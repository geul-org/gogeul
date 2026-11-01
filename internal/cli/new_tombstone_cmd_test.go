//ff:func feature=cli type=command control=sequence
//ff:what TestNewTombstoneCmd — a version is deleted with a reason; --reason is required; a bad address fails
package cli

import (
	"strings"
	"testing"
)

func TestNewTombstoneCmd(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	h.must("tombstone", addr+"@1", "--reason", "2")
	if out := h.must("show", "versions", addr); !strings.Contains(out, "deleted") {
		t.Fatal(out)
	}
	if _, err := h.run("", "tombstone", addr+"@2"); err == nil {
		t.Fatal("no reason: want error")
	}
	if _, err := h.run("", "tombstone", "x@2", "--reason", "1"); err == nil {
		t.Fatal("bad address: want error")
	}
}
