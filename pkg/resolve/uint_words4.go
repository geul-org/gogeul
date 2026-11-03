//ff:func feature=resolve type=helper control=sequence
//ff:what uintWords4 — a 64-bit document ID as its 4 words, high first
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// uintWords4 writes doc as the 4-word document ID.
func uintWords4(doc uint64) []core.Word {
	return []core.Word{core.Word(doc >> 48), core.Word(doc >> 32), core.Word(doc >> 16), core.Word(doc)}
}
