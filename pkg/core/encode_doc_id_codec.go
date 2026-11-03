//ff:func feature=core type=encoder control=sequence
//ff:what EncodeDocIDCodec — Meta DOC-ID of a codec document: one CODEC-POINTER(Issuer, codec ID, version 0) Reference
package core

import "fmt"

// EncodeDocIDCodec builds the Meta DOC-ID of a codec document: its payload is
// the Reference CODEC-POINTER(issuer, codec, 0). The version is 0 because the
// version list, not the document, assigns versions.
func EncodeDocIDCodec(issuer []Word, codec uint32) (Record, error) {
	ptr, err := EncodeCodecPointer(issuer, codec, 0)
	if err != nil {
		return Record{}, fmt.Errorf("geul/core: DOC-ID: %w", err)
	}
	ref, _ := EncodeRecord(ptr)
	return Record{Kind: KindMeta, Codec: MetaDocID, Payload: ref}, nil
}
