//ff:func feature=schema type=rule control=sequence
//ff:what TestListDoc — the first PUBLICATION's document; a non-PUBLICATION first record gives ""
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestListDoc(t *testing.T) {
	if d := listDoc([]rec{recOf(edgeOf(core.EdgePublication, 1, append([]core.Word{0x0418, 0, 0, 0xffff, 0, 0, 2, 0, 1}, make([]core.Word, 16)...)...))}); d != "0xFFFF000000000002" {
		t.Fatalf("got %q", d)
	}
	if d := listDoc([]rec{recOf(edgeOf(0, 1))}); d != "" {
		t.Fatalf("bad: %q", d)
	}
}
