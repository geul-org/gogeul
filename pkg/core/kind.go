//ff:type feature=core type=model
//ff:what Kind — the record kind of a GEUL TAG WORD (Reference/Value/Edge/Meta), read from its section
package core

// Kind is the record kind of a GEUL TAG WORD `000K KKKC CCCC LLLL`. Its value
// is a logical kind: KKKK is the section (w>>9, 32 codes each), C the 5-bit
// offset in it, and the Kind is a run of sections — Reference 0~2
// (w < 0x0600), Value 3·4 (w < 0x0A00), Edge 5~13 (w < 0x1C00), Meta 14·15
// (KindOf). A decoder tells
// the four kinds apart from the TAG WORD alone, without a codebook.
type Kind uint8

const (
	// KindReference is a record that points at a target (sections 0~2).
	KindReference Kind = 0
	// KindValue is a record that carries a value (sections 3·4).
	KindValue Kind = 1
	// KindEdge is a record that composes other records (sections 5~13).
	KindEdge Kind = 2
	// KindMeta is a record about the document (sections 14·15).
	KindMeta Kind = 3
)
