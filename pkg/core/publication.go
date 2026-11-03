//ff:type feature=core type=model
//ff:what Publication — one version list record: PUBLICATION, DEPRECATION or TOMBSTONE
package core

// Publication is one version list record (GEUL t002 docs/edge/PUBLICATION).
// Codec says which: EdgePublication carries Hash, EdgeTombstone carries Reason,
// EdgeDeprecation carries neither. Issuer (2/4/16 words), Doc (the 4-word
// document ID) and Version are the record's DOC-POINTER: one version of the
// document the version list (`<Issuer>/<문서 ID>/versions.geul`) lists.
type Publication struct {
	Codec   uint16
	EdgeID  uint64
	Issuer  []Word
	Doc     []Word
	Version uint32
	Hash    [32]byte
	Reason  uint16
}
