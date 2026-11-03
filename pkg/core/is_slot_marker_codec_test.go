//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsSlotMarkerCodec — only C 0, 1, 2 are slot marker codes; C 3~5 (pointers, LOCAL) are not
package core

import "testing"

func TestIsSlotMarkerCodec(t *testing.T) {
	for c := uint16(0); c <= 0x3F; c++ {
		if got := IsSlotMarkerCodec(c); got != (c <= 2) {
			t.Fatalf("C=0x%02X: %v", c, got)
		}
	}
}
