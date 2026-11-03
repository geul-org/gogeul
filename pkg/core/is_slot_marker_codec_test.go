//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsSlotMarkerCodec — only C 0, 1, 2 are slot marker codes; the empty C 3·4, LOCAL and the rest of the 96 Reference codes are not
package core

import "testing"

func TestIsSlotMarkerCodec(t *testing.T) {
	for c := uint16(0); c <= 0x5F; c++ {
		if got := IsSlotMarkerCodec(c); got != (c <= 2) {
			t.Fatalf("C=0x%02X: %v", c, got)
		}
	}
}
