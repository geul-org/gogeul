//ff:func feature=core type=helper control=sequence
//ff:what IsQuantified — whether a Reference C is one of the 5Q scheme codes (section 0 offset 5~29, section 1 offset 0~29)
package core

// IsQuantified reports whether the Reference C belongs to a 5Q scheme
// (refSlot): C 0x05~0x1D (entity schemes) or 0x20~0x3D (EXTERNAL schemes).
// Slot markers, the empty codes and section 2 (pointers, REFERENCE-ISSUER)
// are outside it.
func IsQuantified(c uint16) bool {
	_, _, ok := refSlot(c)
	return ok
}
