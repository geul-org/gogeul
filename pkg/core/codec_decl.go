//ff:type feature=core type=model
//ff:what CodecDecl — a CODEC Edge: Edge Id, the FIELD it belongs to, the Kind·Codec target word
package core

// CodecDecl is a CODEC Edge (GEUL t002 docs/edge/CODEC): one Kind·Codec the
// FIELD Field takes, as a target word: TargetWord (LLLL 0000, a Reference's
// Quantifier kept) or a section wildcard with LLLL 1111 (CodecMatches).
type CodecDecl struct {
	EdgeID uint64
	Field  uint64
	Target Word
}
