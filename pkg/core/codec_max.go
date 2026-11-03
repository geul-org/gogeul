//ff:func feature=core type=helper control=selection
//ff:what CodecMax — the largest Codec a Kind holds (Reference 0x5F, Value and Meta 0x3F, Edge 0x11F)
package core

// CodecMax returns the largest Record Codec of k: 0x5F for Reference (sections
// 0~2, 96 codes), 0x3F for Value and Meta (two sections, 64 codes each) and
// 0x11F for Edge (sections 5~13, 288 codes, Codec = code - 0xA0).
func CodecMax(k Kind) uint16 {
	switch k {
	case KindReference:
		return 0x5F
	case KindEdge:
		return 0x11F
	default:
		return 0x3F
	}
}
