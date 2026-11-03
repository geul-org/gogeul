//ff:func feature=core type=decoder control=sequence
//ff:what DecodeCodecPointer — the Issuer, codec ID and version ID of a CODEC-POINTER Reference
package core

import "fmt"

// DecodeCodecPointer splits r, which must be a CODEC-POINTER Reference
// (0x03CL), into its Issuer (the payload less 4 words: a canonical 2/4/16-word
// Issuer, so a Payload of 6, 8 or 20 words), its 2-word codec ID and its 2-word
// version ID (0 is the latest).
func DecodeCodecPointer(r Record) (issuer []Word, codec, version uint32, err error) {
	w := len(r.Payload) - 4
	if r.Kind != KindReference || r.Codec != CodecCodecPointer || w < 2 || !IssuerCanonical(r.Payload[:w]) {
		return nil, 0, 0, fmt.Errorf("geul/core: expected CODEC-POINTER(Issuer, codec ID, version ID), got kind %d codec 0x%02X, %d words", r.Kind, r.Codec, len(r.Payload))
	}
	p := r.Payload[w:]
	return r.Payload[:w], uint32(p[0])<<16 | uint32(p[1]), uint32(p[2])<<16 | uint32(p[3]), nil
}
