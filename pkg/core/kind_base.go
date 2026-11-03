//ff:func feature=core type=helper control=selection
//ff:what kindBase — the first TAG WORD of a Kind's zone (Reference 0x0000, Value 0x0400, Edge 0x0800, Meta 0x1C00)
package core

// kindBase returns the TAG WORD where k's zone starts: Reference 0x000L, Value
// 0x040L, Edge 0x080L, Meta 0x1C0L. A TAG WORD is kindBase(k) + Codec<<4 + L.
func kindBase(k Kind) Word {
	switch k {
	case KindReference:
		return 0x0000
	case KindValue:
		return 0x0400
	case KindEdge:
		return 0x0800
	default:
		return 0x1C00
	}
}
