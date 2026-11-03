//ff:func feature=core type=helper control=selection
//ff:what CodecMatches — whether a record's Kind·Codec fits a CODEC target (exact or Kind wildcard)
package core

// CodecMatches reports whether a record of Kind k and Codec codec fits the
// CODEC target t (GEUL edge/CODEC). With LLLL 0000 the Kind and Codec must be
// equal; this is the only way a slot marker (0x0000 ABSENT, 0x0010
// UNKNOWN-VALUE, 0x0020 NO-VALUE) is accepted. With LLLL 1111 the C bits are
// not read: any record of the target's Kind (KKK) matches, except the slot
// markers. Several Quantifiers or codes are taken by listing several CODECs.
func CodecMatches(t Word, k Kind, codec uint16) bool {
	tk, tc := SplitTarget(t)
	switch {
	case tk != k:
		return false
	case t&0xF == 0xF:
		return k != KindReference || !IsSlotMarkerCodec(codec)
	default:
		return tc == codec
	}
}
