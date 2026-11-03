//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestInvPath — P entries through the Issuer table, named through the path table; bad numbers fail
package pack

import (
	"errors"
	"testing"
)

func TestInvPath(t *testing.T) {
	inv := inventory{issuers: []string{"0x00000000"}, paths: []string{"issuers/0x00000001.geul"}}
	if p, err := invPath(inv, invEntry{kind: invPDel, doc: 8}); err != nil || p != "0x00000000/docs/0x0000000000000008/versions.geul" {
		t.Fatal(p, err)
	}
	if p, err := invPath(inv, invEntry{kind: invNamed}); err != nil || p != "issuers/0x00000001.geul" {
		t.Fatal(p, err)
	}
	for _, e := range []invEntry{{kind: invP, ref: 1}, {kind: invNamedDel, ref: 1}} {
		if _, err := invPath(inv, e); !errors.Is(err, ErrCorrupt) {
			t.Fatal(e, err)
		}
	}
}
