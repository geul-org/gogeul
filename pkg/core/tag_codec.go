//ff:func feature=core type=decoder control=sequence
//ff:what tagCodec — the Kind and Record Codec of a TAG or target WORD (LLLL ignored)
package core

// tagCodec splits w into its Kind (KindOf) and the Record Codec: the 9-bit
// code KKKKC (w>>4) less the start of the Kind's zone (kindBase), so Reference
// gives 0x00~0x5F, Value and Meta 0x00~0x3F and Edge code - 0xA0. LLLL is not
// read.
func tagCodec(w Word) (Kind, uint16) {
	k := KindOf(w)
	return k, uint16(w&0x1FF0-kindBase(k)) >> 4
}
