//ff:func feature=core type=encoder control=sequence
//ff:what EncodeIssuerPointer — the ISSUER-POINTER Reference of a canonical 2/4/16-word Issuer
package core

import "fmt"

// EncodeIssuerPointer builds the ISSUER-POINTER Reference (0x040L, Identified
// only) for a canonical Issuer of 2, 4 or 16 words (IssuerCanonical).
func EncodeIssuerPointer(issuer []Word) (Record, error) {
	if !IssuerCanonical(issuer) {
		return Record{}, fmt.Errorf("geul/core: ISSUER-POINTER needs a canonical 2/4/16-word Issuer (got %d words)", len(issuer))
	}
	return Record{Kind: KindReference, Codec: CodecIssuerPointer, Payload: append([]Word{}, issuer...)}, nil
}
