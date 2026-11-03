//ff:func feature=core type=helper control=sequence
//ff:what validTarget — whether a word can be a CODEC target (prefix 000, LLLL 0000 or 1111)
package core

// validTarget reports whether t can be a CODEC target (edge/CODEC): the top 3
// bits are 0 and LLLL is 0000 (exact) or 1111 (any record of the Kind). C is
// not read for a wildcard; C 0 is only recommended.
func validTarget(t Word) bool {
	return t&0xE000 == 0 && (t&0xF == 0 || t&0xF == 0xF)
}
