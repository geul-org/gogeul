//ff:func feature=resolve type=helper control=sequence
//ff:what issuerOfID — an Issuer segment that must be an ID (8, 16 or 64 hex digits), in canonical width; else ErrNotFound
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// issuerOfID reads s as an Issuer ID of 8, 16 or 64 lowercase hex digits and
// returns it in canonical width (normalIssuer). A name or any other spelling is
// ErrNotFound.
func issuerOfID(s string) ([]core.Word, error) {
	c, ok := normalIssuer(s)
	if !ok {
		return nil, fmt.Errorf("%w: %q is not an Issuer ID", ErrNotFound, s)
	}
	b, _ := parseHexID(c, (len(c)-2)/4)
	return hexWords(b), nil
}
