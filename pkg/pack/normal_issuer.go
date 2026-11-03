//ff:func feature=pack type=helper control=sequence
//ff:what normalIssuer — an Issuer ID of 8, 16 or 64 hex digits as its canonical hex (zero padding dropped)
package pack

import "github.com/geul-org/gogeul/pkg/core"

// normalIssuer reads s, `0x` + 8, 16 or 64 lowercase hex digits, and returns
// the canonical hex of that Issuer (core.IssuerNormal), the form the Issuer
// table and the stored paths use.
func normalIssuer(s string) (string, bool) {
	n := len(s) - 2
	if n != 8 && n != 16 && n != 64 {
		return "", false
	}
	b, ok := parseHex(s, n)
	if !ok {
		return "", false
	}
	ws, _ := core.ReadWords(b)
	c, _ := core.IssuerNormal(ws)
	return core.FormatID(c), true
}
