//ff:func feature=schema type=rule control=sequence
//ff:what checkEmbedded — M4~M6: a Meta carries one Value of an accepted codec that passes the Value rules
package schema

import "fmt"

// checkEmbedded applies rule to Meta m: its payload is one Value whose codec
// accept allows, and that Value passes the V rules (M6).
func checkEmbedded(m rec, rule, expected string, accept func(codec uint16) bool) []Fact {
	v, ok := embeddedValue(m)
	if !ok {
		return []Fact{factAt(rule, Error, m, "Payload", "Value 레코드 하나", fmtWords(len(m.R.Payload)))}
	}
	if !accept(v.R.Codec) {
		return []Fact{factAt(rule, Error, m, "내장 Value", expected, fmt.Sprintf("Value 0x%02X", v.R.Codec))}
	}
	return checkValue(v)
}
