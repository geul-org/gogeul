//ff:func feature=core type=helper control=selection
//ff:what refSlot — the 5Q scheme and Quantifier of a Reference Codec from its section KKKK and offset C, the only 5Q test
package core

// refSlot reads a Reference Codec as a 5Q code. In TAG `000K KKKC CCCC LLLL`
// the Codec is the section KKKK (codec>>5) and the 5-bit offset C (codec &
// 0x1F): in section 0 the entity schemes 1~5 sit at 5 <= C < 30 (scheme C/5,
// Quantifier C%5), in section 1 the EXTERNAL schemes 6~11 at C < 30 (scheme
// 6 + C/5, Quantifier C%5). The bool is false for every other Codec — slot
// markers, the empty codes, section 2 (ISSUER-POINTER, DOC-POINTER,
// CODEC-POINTER, REFERENCE-ISSUER) — where a Quantifier would be misread.
func refSlot(codec uint16) (scheme, q uint16, ok bool) {
	c := codec & 0x1F
	switch {
	case codec>>5 == 0 && c >= 5 && c < 30:
		return c / 5, c % 5, true
	case codec>>5 == 1 && c < 30:
		return 6 + c/5, c % 5, true
	default:
		return 0, 0, false
	}
}
