//ff:func feature=cli type=command control=sequence
//ff:what TestNewShowCmdVersions — show versions lists a document's versions; `versions` alone is an ordinary argument and extra arguments fail
package cli

import (
	"strings"
	"testing"
)

func TestNewShowCmdVersions(t *testing.T) {
	h := newHarness(t)
	addr, _ := h.doc()
	if out := h.must("show", "versions", addr); !strings.HasPrefix(out, "1\t") || !strings.HasSuffix(out, "NEXT-LOCAL-ID 4\n") {
		t.Fatal(out)
	}
	if _, err := h.run("", "show", "versions"); err == nil {
		t.Fatal("no such file or address: want error")
	}
	if _, err := h.run("", "show", "versions", addr, "more"); err == nil {
		t.Fatal("three arguments: want error")
	}
	if _, err := h.run("", "show", addr, addr); err == nil {
		t.Fatal("two documents: want error")
	}
	if _, err := h.run("", "log", addr); err == nil {
		t.Fatal("log is gone: want error")
	}
}
