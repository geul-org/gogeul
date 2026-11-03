//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefCodec — C = scheme×5 + quantifier, the inverse of RefScheme and RefQuantifier over every 5Q code
package core

import "testing"

func TestRefCodec(t *testing.T) {
	if RefCodec(1, 0) != 0x05 || RefCodec(11, 4) != 0x3B {
		t.Fatal("ends")
	}
	for c := uint16(0x05); c <= 0x3B; c++ {
		s, _ := RefScheme(c)
		q, _ := RefQuantifier(c)
		if RefCodec(s, q) != c {
			t.Fatalf("C %X", c)
		}
	}
}
