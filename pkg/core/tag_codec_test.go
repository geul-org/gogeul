//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestTagCodec — Kind and zone Codec of a TAG WORD, Length bits ignored: first and last code of each zone
package core

import "testing"

func TestTagCodec(t *testing.T) {
	cases := []struct {
		w Word
		k Kind
		c uint16
	}{
		{0x0000, KindReference, 0}, {0x05FF, KindReference, 0x5F},
		{0x0601, KindValue, 0}, {0x09F0, KindValue, 0x3F},
		{0x0A02, KindEdge, 0}, {0x1BFF, KindEdge, 0x11F},
		{0x1C0E, KindMeta, 0}, {0x1FF1, KindMeta, 0x3F},
	}
	for _, c := range cases {
		if k, codec := tagCodec(c.w); k != c.k || codec != c.c {
			t.Fatalf("%04X: %v %X, want %v %X", c.w, k, codec, c.k, c.c)
		}
	}
}
