//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestKindOf — KKK of 000K KKCC CCCC LLLL: 000 Reference, 001 Value, 010~110 Edge, 111 Meta (zone boundary WORDs of migration.md §7.3)
package core

import "testing"

func TestKindOf(t *testing.T) {
	want := map[Word]Kind{
		0x0000: KindReference, 0x03FF: KindReference,
		0x0400: KindValue, 0x07FF: KindValue,
		0x0800: KindEdge, 0x0BFF: KindEdge, 0x0C00: KindEdge, 0x1BFF: KindEdge,
		0x1C00: KindMeta, 0x1FFF: KindMeta,
	}
	for w, k := range want {
		if got := KindOf(w); got != k {
			t.Fatalf("%04X: %v, want %v", w, got, k)
		}
	}
}
