//ff:func feature=cli type=command control=sequence
//ff:what TestNewNameCmd — a name binds and resolves; a version, a bad address and no data folder fail
package cli

import (
	"strings"
	"testing"
)

func TestNewNameCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	h.must("name", addr, "adams")
	if out := h.must("show", "me/docs/adams/latest"); !strings.Contains(out, "ALIAS") {
		t.Fatal(out)
	}
	if _, err := h.run("", "name", addr+"@1", "z"); err == nil {
		t.Fatal("a version: want error")
	}
	if _, err := h.run("", "name", "x", "y"); err == nil {
		t.Fatal("bad address: want error")
	}
	h.cwd = t.TempDir()
	if _, err := h.run("", "name", addr, "y"); err == nil {
		t.Fatal("no data folder: want error")
	}
}
