//ff:func feature=cli type=command control=sequence
//ff:what TestAppLs — text and JSON; no data folder, an unknown Issuer and a missing archive fail
package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAppLs(t *testing.T) {
	h := newHarness(t)
	h.doc()
	if out := h.must("ls", "me"); !strings.HasPrefix(out, "0x0000000000000001") {
		t.Fatal(out)
	}
	if out := h.must("ls", "--format", "json"); !strings.HasPrefix(out, "[") {
		t.Fatal(out)
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "ls"); err == nil {
		t.Fatal("no data folder: want error")
	}
	h = newHarness(t)
	if _, err := h.run("", "ls", "nobody"); err == nil {
		t.Fatal("unknown Issuer: want error")
	}
	os.RemoveAll(filepath.Join(h.data, ".well-known"))
	if _, err := h.run("", "ls", "--data", h.data); err == nil {
		t.Fatal("no archive: want error")
	}
}
