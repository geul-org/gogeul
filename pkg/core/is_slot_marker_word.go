//ff:func feature=core type=helper control=sequence
//ff:what IsSlotMarkerWord — whether a WORD is 0x0000, 0x0010 or 0x0020
package core

// IsSlotMarkerWord reports whether w is 0x0000, 0x0010 or 0x0020: as a TAG
// WORD, a one-word slot marker; as a CODEC target, the exact target naming
// that marker. No codebook is needed.
func IsSlotMarkerWord(w Word) bool {
	return w&0xF == 0 && w <= NoValue
}
