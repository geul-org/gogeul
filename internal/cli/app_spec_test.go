//ff:func feature=cli type=helper control=sequence
//ff:what TestAppSpec — empty without a data folder or when a spec document does not read
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/internal/config"
)

func TestAppSpec(t *testing.T) {
	a := &app{env: config.Env{Getenv: func(string) string { return "" }, Cwd: t.TempDir(), ConfigDir: t.TempDir()}}
	a.cfg, _ = config.Load(a.env)
	if len(a.spec()) != 0 {
		t.Fatal("no data folder: want empty")
	}
	h := newHarness(t)
	h.tripleSpec()
	a.env.Cwd = h.data
	a.cfg, _ = config.Load(a.env)
	if len(a.spec()) != 1 {
		t.Fatal("TRIPLE: want one")
	}
	list := filepath.Join(h.data, ".well-known", "geul", "0x00000000", "codecs", "0x00000AA0", "versions.geul")
	os.WriteFile(list, []byte{1}, 0o644)
	if len(a.spec()) != 0 {
		t.Fatal("broken spec: want empty")
	}
}
