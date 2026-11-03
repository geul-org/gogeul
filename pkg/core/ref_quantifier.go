//ff:func feature=core type=helper control=sequence
//ff:what RefQuantifier — the Quantifier C%5 of a quantified Reference C
package core

// RefQuantifier returns the Quantifier of the Reference C (C%5). The bool is
// false outside 5 <= C < 0x3C (IsQuantified): a slot marker or pointer number
// read with C%5 would look like Generic~Variable.
func RefQuantifier(c uint16) (uint16, bool) {
	if !IsQuantified(c) {
		return 0, false
	}
	return c % 5, true
}
