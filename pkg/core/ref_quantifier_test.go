//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefQuantifier — the Quantifier is the C % 5 (C the offset in the section) (C 0x20 is Identified, not C%5 = Universal); false outside the 5Q codes
package core

import "testing"

func TestRefQuantifier(t *testing.T) {
	for c, want := range map[uint16]uint16{0x05: 0, 0x06: 1, 0x09: 4, 0x1D: 4, 0x20: 0, 0x23: 3, 0x24: 4, 0x25: 0, 0x39: 0, 0x3D: 4} {
		if q, ok := RefQuantifier(c); !ok || q != want {
			t.Fatalf("C %X: %d %v", c, q, ok)
		}
	}
	for _, c := range []uint16{2, 3, 4, 0x1E, 0x1F, 0x3E, 0x3F, 0x40, 0x42, 0x5D, 0x5F} {
		if _, ok := RefQuantifier(c); ok {
			t.Fatalf("C %X: want false", c)
		}
	}
}
