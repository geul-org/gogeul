//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestAddrOfPath — each form gives its Kind; the gone docs/·codecs/ lists, short, unknown-folder, named and bad-Issuer paths fail
package resolve

import (
	"errors"
	"strings"
	"testing"
)

func TestAddrOfPath(t *testing.T) {
	h := "0x" + strings.Repeat("ab", 32) + ".geul"
	for p, k := range map[string]AddrKind{
		"issuers/latest.geul":                            AddrIssuers,
		"0x00000000/codecs/0x00000b80/latest.geul":       AddrCodec,
		"0x00000000/docs/0x0000000000000007/latest.geul": AddrDoc,
		"0x00000000/hashes/" + h:                         AddrHash,
	} {
		if a, err := addrOfPath(p, issuerOfID); err != nil || a.Kind != k {
			t.Fatal(p, a, err)
		}
	}
	for _, p := range []string{"issuers", "0x00000000/docs", "0x00000000/other/latest.geul", "0x00000000/docs/0x0000000000000007/x/latest.geul", "geul.org/docs/latest.geul",
		"0x00000000/codecs/versions.geul", "0x00000000/codecs/latest.geul", "0x00000000/docs/0x00000001.geul", "0x00000000/docs/versions.geul", "0x00000000/docs/seven/latest.geul"} {
		if _, err := addrOfPath(p, issuerOfID); !errors.Is(err, ErrNotFound) {
			t.Fatal(p, err)
		}
	}
}
