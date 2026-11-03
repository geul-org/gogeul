//ff:func feature=resolve type=helper control=sequence
//ff:what docOfID — a document segment that must be a 4-word ID (16 hex digits); else ErrNotFound
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// docOfID reads s as a 4-word document ID (16 lowercase hex digits). issuer is
// not used: an address names documents by ID only.
func docOfID(_ []core.Word, s string) ([]core.Word, error) {
	b, ok := parseHexID(s, 4)
	if !ok {
		return nil, fmt.Errorf("%w: %q is not a document ID", ErrNotFound, s)
	}
	return hexWords(b), nil
}
