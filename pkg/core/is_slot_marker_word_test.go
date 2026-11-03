//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsSlotMarkerWord — over 0x0000~0x1FFF only 0x0000, 0x0010, 0x0020 are slot marker WORDs
package core

import "testing"

func TestIsSlotMarkerWord(t *testing.T) {
	for w := Word(0); w < 0x2000; w++ {
		want := w == 0x0000 || w == 0x0010 || w == 0x0020
		if got := IsSlotMarkerWord(w); got != want {
			t.Fatalf("0x%04X: %v", w, got)
		}
	}
}
