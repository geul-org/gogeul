//ff:func feature=core type=helper control=sequence
//ff:what RefCodec — the Reference C of a scheme and Quantifier (section 0 scheme*5 + q, section 1 0x20 + (scheme-6)*5 + q)
package core

// RefCodec returns the Reference C of scheme (1~11) with Quantifier q (0~4):
// scheme*5 + q for the entity schemes 1~5 (section 0), 0x20 + (scheme-6)*5 + q
// for the EXTERNAL schemes 6~11 (section 1). The caller keeps both in range.
func RefCodec(scheme, q uint16) uint16 {
	if scheme < SchemeExternalPointer1Word {
		return scheme*5 + q
	}
	return 0x20 + (scheme-SchemeExternalPointer1Word)*5 + q
}
