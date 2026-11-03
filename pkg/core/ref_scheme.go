//ff:func feature=core type=helper control=sequence
//ff:what RefScheme — the scheme number (1~11) of a quantified Reference C
package core

// RefScheme returns the scheme of the Reference Codec (1~11, refSlot): C/5 in
// section 0, 6 + C/5 in section 1, C the 5-bit offset in the section. The bool is false outside
// the 5Q codes (IsQuantified).
func RefScheme(c uint16) (uint16, bool) {
	s, _, ok := refSlot(c)
	return s, ok
}
