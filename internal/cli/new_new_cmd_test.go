//ff:func feature=cli type=command control=sequence
//ff:what TestNewNewCmd — a draft with DOC-ID and ISSUER-KEY; GEULISSUER fills a missing Issuer; -o is required
package cli

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNewNewCmd(t *testing.T) {
	h := newHarness(t)
	h.must("issuer", "new", "--name", "me")
	f := filepath.Join(h.data, "n.geul")
	h.must("new", "me/0x0000000000000003", "-o", f)
	if out := h.must("show", f); !strings.Contains(out, "ISSUER-KEY") || !strings.Contains(out, "/0x0000000000000003") {
		t.Fatal(out)
	}
	h.env["GEULISSUER"] = "0x00010002"
	g := filepath.Join(h.data, "g.geul")
	h.must("new", "0x0000000000000004", "-o", g)
	if out := h.must("show", g); !strings.Contains(out, "DOC-ID 0x00010002/0x0000000000000004") || strings.Contains(out, "ISSUER-KEY") {
		t.Fatal(out)
	}
	if _, err := h.run("", "new", "me/0x0000000000000005"); err == nil {
		t.Fatal("no -o: want error")
	}
}
