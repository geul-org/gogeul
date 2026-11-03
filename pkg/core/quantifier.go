//ff:type feature=core type=model
//ff:what Quantifier constants — the 5Q order of a Reference scheme (C % 5)
package core

// The five Quantifiers of a Reference scheme, in 5Q order: the Quantifier of
// a quantified Reference is C % 5 (RefQuantifier).
const (
	// QuantIdentified names one fixed target.
	QuantIdentified uint16 = 0
	// QuantGeneric names a kind in general.
	QuantGeneric uint16 = 1
	// QuantFreeChoice names any one target.
	QuantFreeChoice uint16 = 2
	// QuantUniversal names every target.
	QuantUniversal uint16 = 3
	// QuantVariable is a variable to be bound.
	QuantVariable uint16 = 4
)
