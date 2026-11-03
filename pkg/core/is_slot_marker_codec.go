//ff:func feature=core type=helper control=sequence
//ff:what IsSlotMarkerCodec — whether a Reference C is a slot marker code (C < 3)
package core

// IsSlotMarkerCodec reports whether the Reference C is one of the slot marker
// codes ABSENT, UNKNOWN-VALUE, NO-VALUE (C < 3). C 3 and 4 are Reserved;
// the pointers are C 0x40~0x42 (grammar reference.md).
func IsSlotMarkerCodec(c uint16) bool {
	return c < slotMarkerCodecs
}
