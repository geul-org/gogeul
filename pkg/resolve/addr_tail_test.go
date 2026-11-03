//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestAddrTail — latest.geul (Latest), versions.geul (Versions) and a nonzero 8-digit version; version 0 and other spellings are bad
package resolve

import (
	"errors"
	"testing"
)

func TestAddrTail(t *testing.T) {
	bad := errors.New("bad")
	if a, err := addrTail(Addr{}, "latest.geul", bad); err != nil || a.Version != Latest || a.Versions {
		t.Fatal(a, err)
	}
	if a, err := addrTail(Addr{}, "versions.geul", bad); err != nil || !a.Versions {
		t.Fatal(a, err)
	}
	if a, err := addrTail(Addr{}, "0x0000000a.geul", bad); err != nil || a.Version != 10 {
		t.Fatal(a, err)
	}
	for _, n := range []string{"0x00000000.geul", "0x0000000A.geul", "0x0000000a", "0x000a.geul", "latest"} {
		if _, err := addrTail(Addr{}, n, bad); err != bad {
			t.Fatal(n, err)
		}
	}
}
