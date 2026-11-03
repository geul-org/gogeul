//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestSplitTarget — the Kind from KKK and the Codec within the Kind zone, LLLL ignored, over the zone edges
package core

import "testing"

func TestSplitTarget(t *testing.T) {
	cases := []struct {
		w Word
		k Kind
		c uint16
	}{
		{0x0000, KindReference, 0}, {0x0060, KindReference, 6}, {0x03FF, KindReference, 0x3F},
		{0x0400, KindValue, 0}, {0x07FF, KindValue, 0x3F},
		{0x0800, KindEdge, 0}, {0x0F9F, KindEdge, 0x79}, {0x1000, KindEdge, 0x80}, {0x1BFF, KindEdge, 0x13F},
		{0x1C00, KindMeta, 0}, {0x1E4F, KindMeta, 0x24}, {0x1FFF, KindMeta, 0x3F},
	}
	for _, c := range cases {
		if k, cc := SplitTarget(c.w); k != c.k || cc != c.c {
			t.Fatalf("%04X: %v %03X", c.w, k, cc)
		}
	}
}
