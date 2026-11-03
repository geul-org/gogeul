//ff:func feature=cli type=command control=sequence
//ff:what TestNewPublishCmd — publish prints the version and hash, --amend replaces the versions; there is no --name
package cli

import (
	"strings"
	"testing"
)

func TestNewPublishCmd(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	h.must("add", draft, "ALIAS", "Q42", `"D"@en`)
	if _, err := h.run("", "publish", draft, "--name", "adams"); err == nil {
		t.Fatal("--name: want error")
	}
	if out := h.must("publish", draft); !strings.HasPrefix(out, "3 0x") {
		t.Fatal(out)
	}
	h.must("publish", draft, "--amend")
	if out := h.must("show", "versions", addr); !strings.HasPrefix(out, "1\t") || strings.Count(out, "\n") != 2 {
		t.Fatal(out)
	}
}
