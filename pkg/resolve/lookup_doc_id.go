//ff:func feature=resolve type=helper control=sequence
//ff:what lookupDocID — the document part of a Lookup address: a 4-word document ID or a 2-word codec ID in canonical hex
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// lookupDocID reads s, the document part of `<Issuer>/<문서>`, as a 4-word
// document ID (16 lowercase hex digits) or a 2-word codec ID (8 digits).
// Anything else is ErrNotFound: documents have no names (§1.1 101). issuer
// only names the Issuer in the error.
func lookupDocID(issuer []core.Word, s string) ([]core.Word, error) {
	if b, ok := parseHexID(s, 4); ok {
		return hexWords(b), nil
	}
	if b, ok := parseHexID(s, 2); ok {
		return hexWords(b), nil
	}
	return nil, fmt.Errorf("%w: %q under %s is not a document or codec ID", ErrNotFound, s, core.FormatID(issuer))
}
