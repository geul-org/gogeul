//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirLookup — Issuer IDs and names, document and codec IDs, with and without a version resolve; document names, malformed addresses and unknown names fail
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirLookup(t *testing.T) {
	d := issuersFixture(t)
	for s, v := range map[string]uint32{"geul.org/0xffff000000000007": 0, "0x00000000/0xffff000000000007@2": 2, "0x0000000000000000/0xffff000000000007": 0, "geul.org/0xffff000000000007@0x00000003": 3} {
		k, err := d.Lookup(s)
		if err != nil || core.FormatID(k.Doc) != "0xffff000000000007" || core.FormatID(k.Issuer) != "0x00000000" || k.Version != v {
			t.Fatalf("%s: %+v %v", s, k, err)
		}
	}
	if k, err := d.Lookup("geul.org/0x00000980@1"); err != nil || core.FormatID(k.Doc) != "0x00000980" || k.Version != 1 {
		t.Fatalf("codec: %+v %v", k, err)
	}
	for _, s := range []string{"seven", "geul.org/seven", "geul.org/0xffff000000000007@x", "nobody/seven", "geul.org/nothing", "0x00000000/0xFFFF000000000007", "0x0000/0xffff000000000007"} {
		if _, err := d.Lookup(s); err == nil {
			t.Fatalf("%s: want error", s)
		}
	}
}
