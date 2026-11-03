//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestAddrPath — a hash file, versions.geul, latest.geul and a version file under the folder
package resolve

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAddrPath(t *testing.T) {
	g := []core.Word{0, 0}
	var h [32]byte
	h[0] = 0xab
	for want, a := range map[string]Addr{
		"0x00000000/hashes/0xAB" + strings.Repeat("0", 62) + ".geul": {Kind: AddrHash, Issuer: g, Hash: h},
		"0x00000000/docs/0x0000000000000007/versions.geul":           {Kind: AddrDoc, Issuer: g, ID: []core.Word{0, 0, 0, 7}, Versions: true},
		"issuers/latest.geul":                          {Kind: AddrIssuers},
		"0x00000000/codecs/0x00000B80/0x00000003.geul": {Kind: AddrCodec, Issuer: g, ID: []core.Word{0, 0x0B80}, Version: 3},
	} {
		if got := a.Path(); got != want {
			t.Fatal(got, want)
		}
	}
}
