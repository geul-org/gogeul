//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestCodecMax — the last Codec of each zone: Edge 0x13F, the others 0x3F
package core

import "testing"

func TestCodecMax(t *testing.T) {
	for k, m := range map[Kind]uint16{KindReference: 0x3F, KindValue: 0x3F, KindEdge: 0x13F, KindMeta: 0x3F} {
		if got := CodecMax(k); got != m {
			t.Fatalf("%v: %X, want %X", k, got, m)
		}
	}
}
