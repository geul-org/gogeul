//ff:func feature=core type=helper control=sequence
//ff:what RefIdentified — fold a quantified Reference C onto its scheme head (Identified)
package core

// RefIdentified folds a quantified Reference C onto the Identified code of its
// scheme (C less its Quantifier, refSlot). Any other C is returned as it is.
func RefIdentified(c uint16) uint16 {
	_, q, ok := refSlot(c)
	if !ok {
		return c
	}
	return c - q
}
