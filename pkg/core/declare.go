//ff:type feature=core type=model
//ff:what Declare — a DECLARE Edge: Edge Id, target Kind·Codec, Issuer and format of an ISSUER target
package core

// Declare is a DECLARE Edge (GEUL t002 docs/edge/DECLARE). Target is the
// declared codec's TAG WORD with only the Kind and Codec bits set. Issuer and
// Format (the 2-word codec number) are set for an ISSUER target. Its fields are the FIELD Edges that
// name this Edge Id, in document order.
type Declare struct {
	EdgeID uint64
	Target Word
	Issuer []Word
	Format uint32
}
