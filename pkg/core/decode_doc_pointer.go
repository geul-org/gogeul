//ff:func feature=core type=decoder control=sequence
//ff:what DecodeDocPointer — the Issuer, document ID and version ID of a DOC-POINTER Reference
package core

import "fmt"

// DecodeDocPointer splits r, which must be a DOC-POINTER Reference (0x004L),
// into its Issuer (the payload less the 6 fixed words, a canonical 2/4/16-word
// Issuer), its 4-word document ID (0 is the Issuer's term definition index
// document) and its 2-word version ID.
func DecodeDocPointer(r Record) (issuer, doc, version []Word, err error) {
	w := len(r.Payload) - 6
	if r.Kind != KindReference || r.Codec != CodecDocPointer || w < 2 || !IssuerCanonical(r.Payload[:w]) {
		return nil, nil, nil, fmt.Errorf("geul/core: expected DOC-POINTER(Issuer, document ID, version ID), got kind %d codec 0x%02X, %d words", r.Kind, r.Codec, len(r.Payload))
	}
	p := r.Payload
	return p[:w], p[w : w+4], p[w+4:], nil
}
