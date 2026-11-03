//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestValidTarget — over the whole 16-bit space a target is valid iff the prefix is 000 and LLLL is 0000 or 1111
package core

import "testing"

func TestValidTarget(t *testing.T) {
	for v := 0; v <= 0xFFFF; v++ {
		w := Word(v)
		want := v < 0x2000 && (v&0xF == 0 || v&0xF == 0xF)
		if got := validTarget(w); got != want {
			t.Fatalf("%04X: %v", v, got)
		}
	}
}
