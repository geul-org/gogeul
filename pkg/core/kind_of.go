//ff:func feature=core type=helper control=sequence
//ff:what KindOf — the Kind of a TAG or target WORD from its KKK bits, the only Kind test
package core

// KindOf reads the Kind of a TAG WORD or CODEC target word from its 3-bit KKK
// `(w>>10)&7`: 0 Reference, 1 Value, 2~6 Edge, 7 Meta. There is no exception;
// every Kind test goes through this table. The prefix bits are not checked.
func KindOf(w Word) Kind {
	return [8]Kind{KindReference, KindValue, KindEdge, KindEdge, KindEdge, KindEdge, KindEdge, KindMeta}[w>>10&7]
}
