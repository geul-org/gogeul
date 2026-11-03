//ff:func feature=schema type=rule control=sequence
//ff:what TestListDoc — the first PUBLICATION's Issuer and document; a non-PUBLICATION first record gives ""
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestListDoc(t *testing.T) {
	if i, d := listDoc([]rec{recOf(edgeOf(core.EdgePublication, 1, append([]core.Word{0x0418, 0, 3, 0xffff, 0, 0, 2, 0, 1}, make([]core.Word, 16)...)...))}); i != "0x00000003" || d != "0xFFFF000000000002" {
		t.Fatalf("got %q %q", i, d)
	}
	if i, d := listDoc([]rec{recOf(edgeOf(0, 1))}); i != "" || d != "" {
		t.Fatalf("bad: %q %q", i, d)
	}
}
