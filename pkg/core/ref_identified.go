//ff:func feature=core type=helper control=sequence
//ff:what RefIdentified — fold a quantified Reference C onto its scheme head (Identified)
package core

// RefIdentified folds a quantified Reference C onto the Identified code of its
// scheme (C - C%5). Any other C is returned as it is.
func RefIdentified(c uint16) uint16 {
	if !IsQuantified(c) {
		return c
	}
	return c - c%5
}
