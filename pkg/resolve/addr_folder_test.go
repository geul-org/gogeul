//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestAddrFolder — the folder of each Kind: issuers, <I>/codecs/<C>, <I>/docs/<D>, <I>/hashes
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAddrFolder(t *testing.T) {
	g := []core.Word{0, 0}
	for want, a := range map[string]Addr{
		"issuers":                            {Kind: AddrIssuers},
		"0x00000000/codecs/0x00000b80":       {Kind: AddrCodec, Issuer: g, ID: []core.Word{0, 0x0B80}},
		"0x00000000/docs/0x0000000000000007": {Kind: AddrDoc, Issuer: g, ID: []core.Word{0, 0, 0, 7}},
		"0x00000000/hashes":                  {Kind: AddrHash, Issuer: g},
	} {
		if got := a.folder(); got != want {
			t.Fatal(got, want)
		}
	}
}
