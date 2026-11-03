//ff:func feature=schema type=rule control=sequence
//ff:what notIdentified — whether a Reference C is a quantified code whose Quantifier is not Identified
package schema

import "github.com/geul-org/gogeul/pkg/core"

// notIdentified reports whether the Reference C belongs to a 5Q scheme (5 <= C
// < 0x3C) with a Quantifier other than Identified. Codes outside the 5Q range
// (slot markers, pointers, REFERENCE-ISSUER) have no Quantifier and are not
// flagged.
func notIdentified(c uint16) bool {
	q, ok := core.RefQuantifier(c)
	return ok && q != core.QuantIdentified
}
