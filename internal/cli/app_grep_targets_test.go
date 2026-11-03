//ff:func feature=cli type=helper control=sequence
//ff:what TestAppGrepTargets — the arguments as given; every document otherwise; no data folder, a missing archive and an unreadable Issuer fail
package cli

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestAppGrepTargets(t *testing.T) {
	h := newHarness(t)
	h.tripleSpec()
	a := &app{}
	if got, _ := a.grepTargets([]string{"x"}); !slices.Equal(got, []string{"x"}) {
		t.Fatal(got)
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "grep", "x"); err == nil {
		t.Fatal("no data folder: want error")
	}
	h.cwd = h.data
	if out := h.must("grep", "DECLARE"); out == "" {
		t.Fatal("every document: want a match")
	}
	os.Chmod(filepath.Join(h.data, ".well-known", "geul", "0x00000000"), 0o000)
	defer os.Chmod(filepath.Join(h.data, ".well-known", "geul", "0x00000000"), 0o755)
	if _, err := h.run("", "grep", "x"); err == nil {
		t.Fatal("unreadable Issuer: want error")
	}
	os.Chmod(filepath.Join(h.data, ".well-known", "geul", "0x00000000"), 0o755)
	os.RemoveAll(filepath.Join(h.data, ".well-known"))
	if _, err := h.run("", "grep", "x", "--data", h.data); err == nil {
		t.Fatal("missing archive: want error")
	}
}
