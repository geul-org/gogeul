//ff:func feature=cli type=helper control=sequence
//ff:what TestAppDir — the data folder as an archive (directory, or pack sets as its Store); none set, or a broken pack archive, fails
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/internal/config"
)

func TestAppDir(t *testing.T) {
	a := &app{env: config.Env{Getenv: func(string) string { return "" }, Cwd: t.TempDir(), ConfigDir: t.TempDir()}}
	a.cfg, _ = config.Load(a.env)
	if _, err := a.dir(); err == nil {
		t.Fatal("none: want error")
	}
	a.env.Flags = map[string]string{"GEULDATA": "/d"}
	a.cfg, _ = config.Load(a.env)
	if d, err := a.dir(); err != nil || d.Root != "/d" || d.Store != nil {
		t.Fatal(err)
	}
	h := newHarness(t)
	h.doc()
	h.packed()
	a = &app{env: config.Env{Getenv: func(string) string { return "" }, Cwd: h.data, ConfigDir: t.TempDir(), Flags: map[string]string{"GEULDATA": h.data}}}
	a.cfg, _ = config.Load(a.env)
	if d, err := a.dir(); err != nil || d.Store == nil {
		t.Fatal("pack", err)
	}
	os.WriteFile(filepath.Join(h.data, ".well-known", "geul", "0x00000000", "latest.txt"), []byte("x"), 0o644)
	a.arch = nil
	if _, err := a.dir(); err == nil {
		t.Fatal("broken pack archive: want error")
	}
}
