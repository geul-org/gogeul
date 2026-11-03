//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestAppCheckArchive — a broken version file is found by schema and by reading it back; a missing archive and an unreadable Issuer fail
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAppCheckArchive(t *testing.T) {
	h := newHarness(t)
	h.cwd = t.TempDir()
	if _, err := h.run("", "check", "--archive"); err == nil {
		t.Fatal("no data folder: want error")
	}
	h.cwd = h.data
	h.doc()
	keys, _ := os.ReadDir(filepath.Join(h.data, ".well-known", "geul"))
	var issuer string
	for _, k := range keys {
		if len(k.Name()) == 66 {
			issuer = k.Name()
		}
	}
	v1 := filepath.Join(h.data, ".well-known", "geul", issuer, "docs", "0x0000000000000001", "0x00000001.geul")
	os.WriteFile(v1, []byte{0, 5}, 0o644)
	out, err := h.run("", "check", "--archive")
	if err == nil || len(out) == 0 {
		t.Fatalf("%q %v", out, err)
	}
	if _, err := h.run("", "check", "--archive", "--data", t.TempDir()); err == nil {
		t.Fatal("no archive: want error")
	}
	os.Chmod(filepath.Join(h.data, ".well-known", "geul", issuer), 0o000)
	defer os.Chmod(filepath.Join(h.data, ".well-known", "geul", issuer), 0o755)
	if _, err := h.run("", "check", "--archive"); err == nil {
		t.Fatal("unreadable Issuer: want error")
	}
}
