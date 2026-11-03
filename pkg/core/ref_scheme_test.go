//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefScheme — scheme = offset/5 in section 0 (LOCAL 0x05 → 1), 6 + offset/5 in section 1 (0x20 → 6, EXTERNAL-ISSUER-16WORD 0x3D → 11); false elsewhere
package core

import "testing"

func TestRefScheme(t *testing.T) {
	for c, want := range map[uint16]uint16{0x05: 1, 0x09: 1, 0x0A: 2, 0x1D: 5, 0x20: 6, 0x25: 7, 0x2A: 8, 0x2F: 9, 0x34: 10, 0x39: 11, 0x3D: 11} {
		if s, ok := RefScheme(c); !ok || s != want {
			t.Fatalf("C %X: %d %v", c, s, ok)
		}
	}
	for _, c := range []uint16{0, 3, 4, 0x1E, 0x1F, 0x3E, 0x3F, 0x40, 0x42, 0x43, 0x5F} {
		if _, ok := RefScheme(c); ok {
			t.Fatalf("C %X: want false", c)
		}
	}
}
