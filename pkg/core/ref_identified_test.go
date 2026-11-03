//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefIdentified — a 5Q code folds to its scheme's Identified code; other codes stay
package core

import "testing"

func TestRefIdentified(t *testing.T) {
	for c, want := range map[uint16]uint16{0x05: 0x05, 0x09: 0x05, 0x3B: 0x37, 0x00: 0x00, 0x04: 0x04, 0x3C: 0x3C, 0x3F: 0x3F} {
		if got := RefIdentified(c); got != want {
			t.Fatalf("C %X: %X, want %X", c, got, want)
		}
	}
}
