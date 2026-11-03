//ff:func feature=core type=encoder control=sequence
//ff:what EncodeCodecPointer — the CODEC-POINTER Reference of (Issuer, codec ID, version ID)
package core

import "fmt"

// EncodeCodecPointer builds the CODEC-POINTER Reference (0x03CL, Identified
// only): a canonical 2/4/16-word Issuer (IssuerCanonical), the 2-word codec ID
// and the 2-word version ID, where version 0 is the latest. The Issuer width
// is the Payload less 4 words, so the Payload is 6, 8 or 20 words.
func EncodeCodecPointer(issuer []Word, codec, version uint32) (Record, error) {
	if !IssuerCanonical(issuer) {
		return Record{}, fmt.Errorf("geul/core: CODEC-POINTER needs a canonical 2/4/16-word Issuer (got %d words)", len(issuer))
	}
	p := append(append([]Word{}, issuer...), Word(codec>>16), Word(codec), Word(version>>16), Word(version))
	return Record{Kind: KindReference, Codec: CodecCodecPointer, Payload: p}, nil
}
