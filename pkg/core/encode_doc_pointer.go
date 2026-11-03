//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDocPointer — the DOC-POINTER Reference of (Issuer, document ID, version ID)
package core

import (
	"fmt"
	"slices"
)

// EncodeDocPointer builds the DOC-POINTER Reference (0x041L, Identified only):
// a canonical 2/4/16-word Issuer (IssuerCanonical), a 4-word document ID (0 is
// the Issuer's term definition index document) and a 2-word version ID, where
// version 0 is the latest.
func EncodeDocPointer(issuer, doc, version []Word) (Record, error) {
	if !IssuerCanonical(issuer) || len(doc) != 4 || len(version) != 2 {
		return Record{}, fmt.Errorf("geul/core: DOC-POINTER needs a canonical 2/4/16-word Issuer, a 4-word document ID and a 2-word version ID (got %d, %d, %d)", len(issuer), len(doc), len(version))
	}
	return Record{Kind: KindReference, Codec: CodecDocPointer, Payload: slices.Concat(issuer, doc, version)}, nil
}
