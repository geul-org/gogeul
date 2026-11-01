//ff:func feature=cli type=command control=sequence
//ff:what TestNewShowCmd — show prints a version as lines, lists versions, and its help gives every form; no argument fails
package cli

import (
	"strings"
	"testing"
)

func TestNewShowCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	if out := h.must("show", addr+"@1"); !strings.Contains(out, "BODY\n#1 LABEL Q42 \"Douglas Adams\"@en\n") {
		t.Fatal(out)
	}
	if out := h.must("show", "versions", addr); !strings.HasPrefix(out, "1\t") {
		t.Fatal(out)
	}
	if out := h.must("show", "--help"); !strings.Contains(out, "geul show versions") {
		t.Fatal(out)
	}
	if _, err := h.run("", "show"); err == nil {
		t.Fatal("want error")
	}
}
