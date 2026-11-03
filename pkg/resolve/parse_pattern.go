//ff:func feature=resolve type=decoder control=sequence
//ff:what parsePattern — one `<Issuer>[/<문서 ID>]` or `*` pattern in canonical hex
package resolve

import (
	"fmt"
	"strings"
)

// parsePattern reads one pattern. Each segment is an ID or `*`; IDs are kept
// canonical (the Issuer at its width, A-F upper case, §1.1 122) to match DocKey.
func parsePattern(s string) (Pattern, error) {
	issuer, doc, _ := strings.Cut(s, "/")
	if doc == "*" {
		doc = ""
	}
	_, docOK := parseHexID(doc, 4)
	if (issuer != "*" && !isIssuerID(issuer)) || (doc != "" && !docOK) {
		return Pattern{}, fmt.Errorf("geul/resolve: bad pattern %q", s)
	}
	if is, ok := normalIssuer(issuer); ok {
		issuer = is
	}
	if strings.HasPrefix(doc, "0x") {
		doc = "0x" + strings.ToUpper(doc[2:])
	}
	return Pattern{Issuer: issuer, Doc: doc}, nil
}
