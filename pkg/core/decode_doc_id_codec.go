//ff:func feature=core type=decoder control=sequence
//ff:what DecodeDocIDCodec — the Issuer and codec ID of a codec document's Meta DOC-ID (CODEC-POINTER)
package core

import "fmt"

// DecodeDocIDCodec reads the Meta DOC-ID of a codec document: exactly one
// CODEC-POINTER Reference (Issuer, codec ID, version ID) with version ID 0. A
// codec document names itself by its codec ID, not by a 4-word document ID.
func DecodeDocIDCodec(r Record) (issuer []Word, codec uint32, err error) {
	if r.Kind != KindMeta || r.Codec != MetaDocID {
		return nil, 0, fmt.Errorf("geul/core: not a DOC-ID (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	ref, n, err := DecodeRecord(r.Payload)
	if err != nil || n != len(r.Payload) {
		return nil, 0, fmt.Errorf("geul/core: DOC-ID payload must be one Reference record")
	}
	issuer, codec, ver, err := DecodeCodecPointer(ref)
	if err != nil || ver != 0 {
		return nil, 0, fmt.Errorf("geul/core: codec DOC-ID must be CODEC-POINTER(Issuer, codec ID, 0)")
	}
	return issuer, codec, nil
}
