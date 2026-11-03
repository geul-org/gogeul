//ff:func feature=core type=helper control=sequence
//ff:what QuantifierName — the Quantifier name of a quantified Reference C, "" outside 5 <= C < 0x3C
package core

// QuantifierName returns the name of the Quantifier of the Reference C (C%5 in
// 5Q order: Identified, Generic, Free-choice, Universal, Variable). Outside
// 5 <= C < 0x3C — slot markers, pointers, CODEC-POINTER, REFERENCE-ISSUER — it
// returns "".
func QuantifierName(c uint16) string {
	q, ok := RefQuantifier(c)
	if !ok {
		return ""
	}
	return [...]string{"Identified", "Generic", "Free-choice", "Universal", "Variable"}[q]
}
