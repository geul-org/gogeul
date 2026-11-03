//ff:func feature=core type=helper control=sequence
//ff:what CodecMax — the largest Codec a Kind holds (0x3F, or 0x13F for the five-KKK Edge zone)
package core

// CodecMax returns the largest Record Codec of k: 0x3F for Reference, Value and
// Meta (64 codes each), 0x13F for Edge, whose zone spans KKK 010~110 (320 codes,
// Codec = code - 0x80).
func CodecMax(k Kind) uint16 {
	if k == KindEdge {
		return 0x13F
	}
	return 0x3F
}
