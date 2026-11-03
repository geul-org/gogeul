//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestExternalPointerWidth — schemes 6, 7, 8 are 1, 2, 4 words; every other scheme is 0
package core

import "testing"

func TestExternalPointerWidth(t *testing.T) {
	want := map[uint16]int{6: 1, 7: 2, 8: 4}
	for s := uint16(0); s < 16; s++ {
		if got := ExternalPointerWidth(s); got != want[s] {
			t.Fatalf("scheme %d: %d want %d", s, got, want[s])
		}
	}
}
