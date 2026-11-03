//ff:func feature=core type=decoder control=sequence
//ff:what tagCodec — the Kind and Record Codec of a TAG or target WORD (LLLL ignored)
package core

// tagCodec splits w into its Kind (KindOf) and the Record Codec: the code w>>4
// less the start of the Kind's zone, so Reference, Value and Meta give the
// 6-bit C and Edge gives code - 0x80. LLLL is not read.
func tagCodec(w Word) (Kind, uint16) {
	k := KindOf(w)
	return k, uint16(w&0x1FF0-kindBase(k)) >> 4
}
