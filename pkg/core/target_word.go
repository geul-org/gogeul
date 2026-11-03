//ff:func feature=core type=helper control=sequence
//ff:what TargetWord — a record's Kind·Codec as a TAG WORD with LLLL 0000
package core

// TargetWord is the TAG WORD of a k/codec record with LLLL = 0000: the DECLARE
// and CODEC target, and the codec ID of a GEUL direct code (its 32-bit value).
func TargetWord(k Kind, codec uint16) Word {
	return tagWord(k, 0, codec)
}
