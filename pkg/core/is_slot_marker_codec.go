//ff:func feature=core type=helper control=sequence
//ff:what IsSlotMarkerCodec — whether a Reference C is a slot marker code (C < 3)
package core

// IsSlotMarkerCodec reports whether the Reference C is one of the slot marker
// codes ABSENT, UNKNOWN-VALUE, NO-VALUE (C < 3). Pointers (C 3, 4) are not.
func IsSlotMarkerCodec(c uint16) bool {
	return c < slotMarkerCodecs
}
