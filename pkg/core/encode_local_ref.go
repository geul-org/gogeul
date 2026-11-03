//ff:func feature=core type=encoder control=sequence
//ff:what EncodeLocalRef — a LOCAL Reference record (Quantifier Identified) for an ID in minimal words
package core

// EncodeLocalRef builds the LOCAL Reference record `0x005L` + id for
// a document-local ID (LOCAL × Identified): Quantifier Identified and
// the target ID in the fewest words. Edge Ids and in-document mentions use it.
func EncodeLocalRef(id uint64) Record {
	return Record{Kind: KindReference, Codec: CodecLocal, Payload: uintWords(id)}
}
