//ff:func feature=core type=helper control=selection
//ff:what kindBase — the first TAG WORD of a Kind's zone (Reference 0x0000, Value 0x0600, Edge 0x0A00, Meta 0x1C00)
package core

// kindBase returns the TAG WORD where k's zone starts: Reference 0x000L, Value
// 0x060L, Edge 0x0A0L, Meta 0x1C0L (section 0, 3, 5, 14). A TAG WORD is
// kindBase(k) + Codec<<4 + L.
func kindBase(k Kind) Word {
	switch k {
	case KindReference:
		return 0x0000
	case KindValue:
		return 0x0600
	case KindEdge:
		return 0x0A00
	default:
		return 0x1C00
	}
}
