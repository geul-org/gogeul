//ff:func feature=core type=helper control=sequence
//ff:what KindOf — the Kind of a TAG or target WORD from its section (w>>9), the only Kind test
package core

// KindOf reads the Kind of a TAG WORD or CODEC target word `000K KKKC CCCC
// LLLL` from its section KKKK `(w>>9)&0xF`: 0~2 Reference (w < 0x0600), 3·4
// Value (w < 0x0A00), 5~13 Edge (w < 0x1C00), 14·15 Meta. There is no exception; every Kind test goes
// through this table. The prefix bits are not checked.
func KindOf(w Word) Kind {
	const r, v, e, m = KindReference, KindValue, KindEdge, KindMeta
	return [16]Kind{r, r, r, v, v, e, e, e, e, e, e, e, e, e, m, m}[w>>9&0xF]
}
