//ff:func feature=core type=helper control=sequence
//ff:what QuantifierName — the Quantifier name of a quantified Reference C, "" outside the 5Q codes
package core

// QuantifierName returns the name of the Quantifier of the Reference C
// (RefQuantifier, in 5Q order: Identified, Generic, Free-choice, Universal, Variable). Outside
// the 5Q codes — slot markers, empty codes, pointers, REFERENCE-ISSUER — it
// returns "".
func QuantifierName(c uint16) string {
	q, ok := RefQuantifier(c)
	if !ok {
		return ""
	}
	return [...]string{"Identified", "Generic", "Free-choice", "Universal", "Variable"}[q]
}
