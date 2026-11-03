//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestAddrDoc — a 4-word document ID is AddrDoc; a name, a 2-word ID and the reserved Issuer list ID are bad; the old docs/codecs list IDs are plain document IDs
package resolve

import (
	"errors"
	"testing"
)

func TestAddrDoc(t *testing.T) {
	bad := errors.New("bad")
	a, err := addrDoc(Addr{}, "0x0000000000000007", "versions.geul", bad)
	if err != nil || a.Kind != AddrDoc || a.ID[3] != 7 || !a.Versions {
		t.Fatal(a, err)
	}
	for _, id := range []string{"seven", "0x00000007", "0xFFFFFFFFFFFFFFFF", "0xffffffffffffffff", "0x0000000000000007/x"} {
		if _, err := addrDoc(Addr{}, id, "latest.geul", bad); err != bad {
			t.Fatal(id, err)
		}
	}
	for _, id := range []string{"0xFFFFFFFFFFFFFFFE", "0xFFFFFFFFFFFFFFFD", "0x0000000000000000"} {
		if a, err := addrDoc(Addr{}, id, "0x00000001.geul", bad); err != nil || a.Kind != AddrDoc || a.Version != 1 {
			t.Fatal(id, a, err)
		}
	}
	if _, err := addrDoc(Addr{}, "0x0000000000000007", "seven.geul", bad); err != bad {
		t.Fatal("bad tail:", err)
	}
}
