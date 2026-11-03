//ff:func feature=core type=encoder control=selection
//ff:what EncodePublication — a PUBLICATION, DEPRECATION or TOMBSTONE Edge from a Publication
package core

import "slices"

// EncodePublication builds the version list record p (GEUL t002
// docs/edge/PUBLICATION): the LOCAL Edge Id, the DOC-POINTER Reference of
// (Issuer, document ID, version ID) and the codec's tail — the content hash
// for PUBLICATION, the reason for TOMBSTONE, nothing for DEPRECATION. p must
// hold a 2/4/16-word Issuer and a 4-word document ID.
func EncodePublication(p Publication) Record {
	id, _ := EncodeRecord(EncodeLocalRef(p.EdgeID))
	ptr, _ := EncodeRecord(Record{Kind: KindReference, Codec: CodecDocPointer, Payload: slices.Concat(p.Issuer, p.Doc, []Word{Word(p.Version >> 16), Word(p.Version)})})
	payload := append(id, ptr...)
	switch p.Codec {
	case EdgePublication:
		h, _ := ReadWords(p.Hash[:])
		payload = append(payload, h...)
	case EdgeTombstone:
		payload = append(payload, Word(p.Reason))
	}
	return Record{Kind: KindEdge, Codec: p.Codec, Payload: payload}
}
