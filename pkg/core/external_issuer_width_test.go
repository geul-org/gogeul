//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestExternalIssuerWidth — schemes 9, 10, 11 are 2, 4, 16 words; every other scheme is 0
package core

import "testing"

func TestExternalIssuerWidth(t *testing.T) {
	want := map[uint16]int{9: 2, 10: 4, 11: 16}
	for s := uint16(0); s < 16; s++ {
		if got := ExternalIssuerWidth(s); got != want[s] {
			t.Fatalf("scheme %d: %d want %d", s, got, want[s])
		}
	}
}
