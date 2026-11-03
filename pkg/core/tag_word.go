//ff:func feature=core type=encoder control=sequence
//ff:what tagWord — pack prefix 000, the Kind zone, the Codec and the 4-bit Length code into a TAG WORD
package core

// tagWord packs a GEUL TAG WORD `000K KKKC CCCC LLLL`: the zone start of k
// (kindBase), the Codec shifted past LLLL and the 4-bit Length code. The caller
// keeps codec within CodecMax(k).
func tagWord(k Kind, lenCode uint8, codec uint16) Word {
	return kindBase(k) + Word(codec)<<4 + Word(lenCode&0xF)
}
