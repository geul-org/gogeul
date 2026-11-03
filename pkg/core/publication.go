//ff:type feature=core type=model
//ff:what Publication — one version list record: PUBLICATION, DEPRECATION or TOMBSTONE
package core

// Publication is one version list record (GEUL t002 docs/edge/PUBLICATION).
// Codec says which: EdgePublication carries Hash, EdgeTombstone carries Reason,
// EdgeDeprecation carries neither. Issuer (2/4/16 words), Doc and Version are
// the record's pointer: Doc is the 4-word document ID of a DOC-POINTER, or the
// 2-word codec ID of a CODEC-POINTER in a codec's version list
// (`<Issuer>/docs/<문서 ID>/versions.geul`, `<Issuer>/codecs/<코덱 ID>/versions.geul`).
type Publication struct {
	Codec   uint16
	EdgeID  uint64
	Issuer  []Word
	Doc     []Word
	Version uint32
	Hash    [32]byte
	Reason  uint16
}
