//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefCodec — C = scheme×5 + q in section 0 and 0x20 + (scheme-6)×5 + q in section 1, the inverse of RefScheme and RefQuantifier over every 5Q code
package core

import "testing"

func TestRefCodec(t *testing.T) {
	ends := map[[2]uint16]uint16{{1, 0}: 0x05, {5, 4}: 0x1D, {6, 0}: 0x20, {6, 4}: 0x24, {8, 4}: 0x2E, {9, 0}: 0x2F, {11, 0}: 0x39, {11, 4}: 0x3D}
	for sq, c := range ends {
		if RefCodec(sq[0], sq[1]) != c {
			t.Fatalf("RefCodec(%d, %d) = 0x%02X, want 0x%02X", sq[0], sq[1], RefCodec(sq[0], sq[1]), c)
		}
	}
	n := 0
	for c := uint16(0); c <= 0x5F; c++ {
		s, ok := RefScheme(c)
		q, _ := RefQuantifier(c)
		if ok && RefCodec(s, q) != c {
			t.Fatalf("C %X", c)
		}
		if ok {
			n++
		}
	}
	if n != 55 {
		t.Fatalf("%d 5Q codes, want 55", n)
	}
}
