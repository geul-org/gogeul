//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefQuantifier — quantifier = C%5 for 5Q codes (0 Identified … 4 Variable); false elsewhere
package core

import "testing"

func TestRefQuantifier(t *testing.T) {
	for c, want := range map[uint16]uint16{0x05: 0, 0x06: 1, 0x09: 4, 0x3B: 4, 0x37: 0} {
		if q, ok := RefQuantifier(c); !ok || q != want {
			t.Fatalf("C %X: %d %v", c, q, ok)
		}
	}
	for _, c := range []uint16{2, 4, 0x3C, 0x3D} {
		if _, ok := RefQuantifier(c); ok {
			t.Fatalf("C %X: want false", c)
		}
	}
}
