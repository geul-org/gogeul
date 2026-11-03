//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.issuerOfName — an Issuer written as an ID (8, 16 or 64 hex digits) or as a name of the Issuer list
package resolve

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// issuerOfName returns the Issuer s names: an ID of 8, 16 or 64 hex digits in
// its canonical width (normalIssuer), otherwise the name looked up in the
// Issuer list (`issuers/`).
func (d Dir) issuerOfName(s string) ([]core.Word, error) {
	if c, ok := normalIssuer(s); ok {
		b, _ := parseHexID(c, (len(c)-2)/4)
		return hexWords(b), nil
	}
	b, err := d.Fetch(context.Background(), issuersPath())
	if err != nil {
		return nil, fmt.Errorf("geul/resolve: Issuer name %q: %w", s, err)
	}
	names, err := ParseIssuerNames(b)
	if err != nil {
		return nil, err
	}
	id, ok := names[s]
	if !ok {
		return nil, fmt.Errorf("%w: Issuer name %q", ErrNotFound, s)
	}
	h, _ := parseHexID(id, (len(id)-2)/4)
	return hexWords(h), nil
}
