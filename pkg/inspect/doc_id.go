//ff:type feature=inspect type=model
//ff:what DocID — Meta DOC-ID: the Issuer and the 4-word document ID (2-word codec ID for a codec document)
package inspect

import "github.com/geul-org/gogeul/pkg/core"

// DocID is Meta DOC-ID: the Issuer (2, 4 or 16 words) and the 4-word document
// ID it issued, or the 2-word codec ID of a codec document (CODEC-POINTER).
type DocID struct {
	Issuer []core.Word
	Doc    []core.Word
}
