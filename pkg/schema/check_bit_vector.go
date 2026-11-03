//ff:func feature=schema type=rule control=sequence
//ff:what checkBitVector — V2: BIT-VECTOR is a 2-word bit count N, then exactly ceil(N/16) words with the unused low bits 0
package schema

import "fmt"

// checkBitVector applies V2 to a BIT-VECTOR Value (grammar value/binary
// bit-vector): the first 2 payload words are the unsigned 32-bit bit count N,
// followed by exactly ceil(N/16) words, MSB-first, whose unused low bits in the
// last word are 0.
func checkBitVector(r rec) []Fact {
	p := r.R.Payload
	if len(p) < 2 {
		return []Fact{factAt("V2", Error, r, "Payload", "Bit Count 2 WORD", fmtWords(len(p)))}
	}
	n := uint64(p[0])<<16 | uint64(p[1])
	if uint64(len(p)-2) != (n+15)/16 {
		return []Fact{factAt("V2", Error, r, "Payload", fmt.Sprintf("2 + ceil(%d/16) WORD", n), fmtWords(len(p)))}
	}
	if n%16 != 0 && uint16(p[len(p)-1])&(1<<(16-n%16)-1) != 0 {
		return []Fact{factAt("V2", Error, r, "Payload", "쓰지 않는 하위 bit 0", fmt.Sprintf("0x%04X", uint16(p[len(p)-1])))}
	}
	return nil
}
