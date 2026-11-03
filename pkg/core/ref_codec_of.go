//ff:func feature=core type=helper control=sequence
//ff:what RefCodec — the Reference C of a scheme and Quantifier (scheme*5 + q)
package core

// RefCodec returns the Reference C of scheme (1~11) with Quantifier q (0~4):
// scheme*5 + q. The caller keeps both in range.
func RefCodec(scheme, q uint16) uint16 {
	return scheme*5 + q
}
