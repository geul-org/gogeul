//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestCodecMax — the last Codec of each zone: Reference 0x5F (three sections), Value and Meta 0x3F, Edge 0x11F (nine sections)
package core

import "testing"

func TestCodecMax(t *testing.T) {
	for k, m := range map[Kind]uint16{KindReference: 0x5F, KindValue: 0x3F, KindEdge: 0x11F, KindMeta: 0x3F} {
		if got := CodecMax(k); got != m {
			t.Fatalf("%v: %X, want %X", k, got, m)
		}
	}
}
