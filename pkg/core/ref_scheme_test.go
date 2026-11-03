//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefScheme — scheme = C/5 for 5Q codes (LOCAL 0x05 → 1, EXTERNAL-ISSUER-16WORD 0x3B → 11); false elsewhere
package core

import "testing"

func TestRefScheme(t *testing.T) {
	for c, want := range map[uint16]uint16{0x05: 1, 0x09: 1, 0x0A: 2, 0x2D: 9, 0x3B: 11} {
		if s, ok := RefScheme(c); !ok || s != want {
			t.Fatalf("C %X: %d %v", c, s, ok)
		}
	}
	for _, c := range []uint16{0, 3, 4, 0x3C, 0x3F} {
		if _, ok := RefScheme(c); ok {
			t.Fatalf("C %X: want false", c)
		}
	}
}
