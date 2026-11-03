//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestListKey — the Issuer list always under 0x00000000 and 0xffffffffffffffff; document and codec lists under their Issuer and 0x…fe / 0x…fd
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestListKey(t *testing.T) {
	is := []core.Word{0, 2}
	for want, a := range map[DocKey]Addr{
		{Issuer: "0x00000000", Doc: "0xffffffffffffffff"}: {Kind: AddrIssuers, Issuer: is},
		{Issuer: "0x00000002", Doc: "0xfffffffffffffffe"}: {Kind: AddrDocs, Issuer: is},
		{Issuer: "0x00000002", Doc: "0xfffffffffffffffd"}: {Kind: AddrCodecs, Issuer: is},
	} {
		if got := listKey(a); got != want {
			t.Fatal(got, want)
		}
	}
}
