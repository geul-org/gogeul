//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerCanonical — 2 words always; 4 words only from 2^32; 16 words only from 2^64; 1 word and other widths fail
package core

import "testing"

func TestIssuerCanonical(t *testing.T) {
	high16 := make([]Word, 16)
	high16[11] = 1
	low16 := make([]Word, 16)
	low16[12] = 1
	for i, c := range []struct {
		issuer []Word
		want   bool
	}{
		{[]Word{0}, false},
		{[]Word{0x1234}, false},
		{[]Word{0, 0}, true},
		{[]Word{0, 1}, true},
		{[]Word{1, 0}, true},
		{[]Word{0xFFFF, 0xFFFF}, true},
		{[]Word{0, 1, 0, 0}, true},
		{[]Word{1, 0, 0, 0}, true},
		{[]Word{0, 0, 0xFFFF, 0xFFFF}, false},
		{[]Word{0, 0, 0, 1}, false},
		{high16, true},
		{low16, false},
		{make([]Word, 16), false},
		{[]Word{1, 2, 3}, false},
		{make([]Word, 8), false},
		{nil, false},
	} {
		if IssuerCanonical(c.issuer) != c.want {
			t.Fatalf("%d: %04X want %v", i, c.issuer, c.want)
		}
	}
}
