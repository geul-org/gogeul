//ff:func feature=core type=helper control=sequence
//ff:what SplitTarget — the Kind and Codec of a DECLARE or CODEC target word
package core

// SplitTarget reads the Kind (section) and Codec of a target word (TargetWord). Its
// LLLL is not read.
func SplitTarget(w Word) (Kind, uint16) {
	return tagCodec(w)
}
