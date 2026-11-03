//ff:func feature=core type=helper control=sequence
//ff:what RefScheme — the scheme number C/5 of a quantified Reference C
package core

// RefScheme returns the scheme of the Reference C (C/5, 1~11). The bool is
// false outside 5 <= C < 0x3C (IsQuantified).
func RefScheme(c uint16) (uint16, bool) {
	if !IsQuantified(c) {
		return 0, false
	}
	return c / 5, true
}
