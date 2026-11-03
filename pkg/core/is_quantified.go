//ff:func feature=core type=helper control=sequence
//ff:what IsQuantified — whether a Reference C is one of the 5Q scheme codes (5 <= C < 0x3C)
package core

// IsQuantified reports whether the Reference C belongs to a 5Q scheme, the
// only range where C/5 is the scheme and C%5 the Quantifier. Slot markers and
// pointers (C < 5), CODEC-POINTER (0x3C) and REFERENCE-ISSUER (0x3D~0x3F) are
// outside it.
func IsQuantified(c uint16) bool {
	return c >= 5 && c < CodecCodecPointer
}
