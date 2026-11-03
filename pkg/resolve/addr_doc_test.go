//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestAddrDoc — a 4-word document ID is AddrDoc; a lookup error passes through; a 2-word ID and the reserved list IDs are bad
package resolve

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAddrDoc(t *testing.T) {
	bad := errors.New("bad")
	a, err := addrDoc(Addr{}, "0x0000000000000007", "versions.geul", docOfID, bad)
	if err != nil || a.Kind != AddrDoc || a.ID[3] != 7 || !a.Versions {
		t.Fatal(a, err)
	}
	if _, err := addrDoc(Addr{}, "seven", "latest.geul", docOfID, bad); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	two := func([]core.Word, string) ([]core.Word, error) { return []core.Word{0, 1}, nil }
	if _, err := addrDoc(Addr{}, "x", "latest.geul", two, bad); err != bad {
		t.Fatal(err)
	}
	for _, id := range []string{"0xffffffffffffffff", "0xfffffffffffffffe", "0xfffffffffffffffd"} {
		if _, err := addrDoc(Addr{}, id, "latest.geul", docOfID, bad); err != bad {
			t.Fatal(id, err)
		}
	}
}
