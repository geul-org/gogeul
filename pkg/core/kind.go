//ff:type feature=core type=model
//ff:what Kind — the record kind of a GEUL TAG WORD (Reference/Value/Edge/Meta), read from KKK
package core

// Kind is the record kind of a GEUL TAG WORD `000K KKCC CCCC LLLL`. Its value is
// a logical kind, not the 3-bit KKK: KKK 000 is Reference, 001 Value, 010~110
// Edge and 111 Meta (KindOf). A decoder tells the four kinds apart from the TAG
// WORD alone, without a codebook.
type Kind uint8

const (
	// KindReference is a record that points at a target (KKK 000).
	KindReference Kind = 0
	// KindValue is a record that carries a value (KKK 001).
	KindValue Kind = 1
	// KindEdge is a record that composes other records (KKK 010~110).
	KindEdge Kind = 2
	// KindMeta is a record about the document (KKK 111).
	KindMeta Kind = 3
)
