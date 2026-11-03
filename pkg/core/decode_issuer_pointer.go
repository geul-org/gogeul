//ff:func feature=core type=decoder control=sequence
//ff:what DecodeIssuerPointer — the Issuer of an ISSUER-POINTER Reference
package core

import "fmt"

// DecodeIssuerPointer returns the Issuer of r, which must be an ISSUER-POINTER
// Reference (0x040L) with a canonical 2/4/16-word Issuer (IssuerCanonical).
func DecodeIssuerPointer(r Record) ([]Word, error) {
	if r.Kind != KindReference || r.Codec != CodecIssuerPointer || !IssuerCanonical(r.Payload) {
		return nil, fmt.Errorf("geul/core: expected ISSUER-POINTER(Issuer 2/4/16 WORD), got kind %d codec 0x%02X, %d words", r.Kind, r.Codec, len(r.Payload))
	}
	return r.Payload, nil
}
