//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestAddrHash — a 64-digit lowercase hash file is AddrHash with its bytes; other names are bad
package resolve

import (
	"errors"
	"strings"
	"testing"
)

func TestAddrHash(t *testing.T) {
	bad := errors.New("bad")
	h := "0x" + strings.Repeat("0a", 32)
	a, err := addrHash(Addr{}, h+".geul", bad)
	if err != nil || a.Kind != AddrHash || a.Hash[0] != 0x0a || a.Hash[31] != 0x0a {
		t.Fatal(a, err)
	}
	for _, n := range []string{h, strings.ToUpper(h) + ".geul", h[:60] + ".geul"} {
		if _, err := addrHash(Addr{}, n, bad); err != bad {
			t.Fatal(n, err)
		}
	}
}
