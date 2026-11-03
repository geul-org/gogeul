//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestSplitTarget — the Kind from the section and the Codec within the Kind zone, LLLL ignored, over the plan2 zone edges
package core

import "testing"

func TestSplitTarget(t *testing.T) {
	cases := []struct {
		w Word
		k Kind
		c uint16
	}{
		{0x0000, KindReference, 0}, {0x0060, KindReference, 6}, {0x0420, KindReference, 0x42}, {0x05FF, KindReference, 0x5F},
		{0x0600, KindValue, 0}, {0x07D0, KindValue, 0x1D}, {0x09DF, KindValue, 0x3D}, {0x09FF, KindValue, 0x3F},
		{0x0A00, KindEdge, 0}, {0x119F, KindEdge, 0x79}, {0x11C0, KindEdge, 0x7C}, {0x1200, KindEdge, 0x80},
		{0x1340, KindEdge, 0x94}, {0x1BD0, KindEdge, 0x11D}, {0x1BFF, KindEdge, 0x11F},
		{0x1C00, KindMeta, 0}, {0x1E4F, KindMeta, 0x24}, {0x1FFF, KindMeta, 0x3F},
	}
	for _, c := range cases {
		if k, cc := SplitTarget(c.w); k != c.k || cc != c.c {
			t.Fatalf("%04X: %v %03X", c.w, k, cc)
		}
	}
}
