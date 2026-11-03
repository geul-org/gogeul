//ff:func feature=resolve type=helper control=sequence
//ff:what normalIssuer — an Issuer ID in 8, 16 or 64 hex digits as its canonical hex (zero padding dropped)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// normalIssuer reads s, `0x` + 8, 16 or 64 hex digits, and returns
// the canonical hex of that Issuer (core.IssuerNormal): `0x0000000000000001`
// and the 64-digit form of the same value are both `0x00000001`.
func normalIssuer(s string) (string, bool) {
	n := (len(s) - 2) / 4
	if n != 2 && n != 4 && n != 16 {
		return "", false
	}
	b, ok := parseHexID(s, n)
	if !ok {
		return "", false
	}
	ws, _ := core.ReadWords(b)
	c, _ := core.IssuerNormal(ws)
	return core.FormatID(c), true
}
