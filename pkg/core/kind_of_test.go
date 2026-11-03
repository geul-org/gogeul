//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestKindOf — the Kind by section (w>>9) at the plan2 boundaries 0x05FF/0x0600, 0x09FF/0x0A00, 0x1BFF/0x1C00; the prefix bits are not read
package core

import "testing"

func TestKindOf(t *testing.T) {
	want := map[Word]Kind{
		0x0000: KindReference, 0x03FF: KindReference, 0x0400: KindReference, 0x05FF: KindReference,
		0x0600: KindValue, 0x07FF: KindValue, 0x0800: KindValue, 0x09FF: KindValue,
		0x0A00: KindEdge, 0x0BFF: KindEdge, 0x0C00: KindEdge, 0x1BFF: KindEdge,
		0x1C00: KindMeta, 0x1FFF: KindMeta,
		0x2000: KindReference, 0xE5FF: KindReference, 0xE600: KindValue, 0xFA00: KindEdge, 0xFC00: KindMeta,
	}
	for w, k := range want {
		if got := KindOf(w); got != k {
			t.Fatalf("%04X: %v, want %v", w, got, k)
		}
	}
}
