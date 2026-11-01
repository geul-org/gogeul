//ff:func feature=cli type=command control=sequence
//ff:what TestNewPublishCmd — publish prints the version and hash, --name binds a name, --amend replaces the versions
package cli

import (
	"strings"
	"testing"
)

func TestNewPublishCmd(t *testing.T) {
	h := newHarness(t)
	addr, draft := h.doc()
	h.must("add", draft, "ALIAS", "Q42", `"D"@en`)
	if out := h.must("publish", draft, "--name", "adams"); !strings.HasPrefix(out, "3 0x") {
		t.Fatal(out)
	}
	if out := h.must("ls", "me"); out != "0x0000000000000001 adams\n" {
		t.Fatal(out)
	}
	h.must("publish", draft, "--amend")
	if out := h.must("show", "versions", addr); !strings.HasPrefix(out, "1\t") || strings.Count(out, "\n") != 2 {
		t.Fatal(out)
	}
}
