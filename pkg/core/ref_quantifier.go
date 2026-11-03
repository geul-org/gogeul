//ff:func feature=core type=helper control=sequence
//ff:what RefQuantifier — the Quantifier (C % 5, C the offset in the section) of a quantified Reference Codec
package core

// RefQuantifier returns the Quantifier of the Reference Codec: C%5 with C the
// 5-bit offset in its section (codec & 0x1F, refSlot). The bool is false
// outside the 5Q codes (IsQuantified): a slot marker or pointer number read
// with %5 would look like Generic~Variable, and %5 over the whole Codec
// misreads section 1 (32 is not a multiple of 5).
func RefQuantifier(c uint16) (uint16, bool) {
	_, q, ok := refSlot(c)
	return q, ok
}
