//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefName — a 5Q code is named by its scheme (any quantifier), the other Reference codes by their own name
package core

import "testing"

func TestRefName(t *testing.T) {
	for c, want := range map[uint16]string{0x09: refName(0x05), 0x3B: refName(0x37), 0x3C: "CODEC-POINTER"} {
		if got := refName(c); got != want || got == "" {
			t.Fatalf("C %X: %q, want %q", c, got, want)
		}
	}
	if refName(0x05) != "LOCAL" {
		t.Fatalf("LOCAL: %q", refName(0x05))
	}
}
