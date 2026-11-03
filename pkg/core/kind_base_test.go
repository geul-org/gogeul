//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestKindBase — zone start TAGs: Reference 0x0000, Value 0x0400, Edge 0x0800, Meta 0x1C00
package core

import "testing"

func TestKindBase(t *testing.T) {
	for k, w := range map[Kind]Word{KindReference: 0x0000, KindValue: 0x0400, KindEdge: 0x0800, KindMeta: 0x1C00} {
		if got := kindBase(k); got != w {
			t.Fatalf("%v: %04X, want %04X", k, got, w)
		}
	}
}
