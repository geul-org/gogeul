//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIssuerNormal — zero-padded 4- and 16-word Issuers reduce to the same canonical Issuer; other widths fail
package core

import (
	"slices"
	"testing"
)

func TestIssuerNormal(t *testing.T) {
	pad := func(n int, tail ...Word) []Word { return append(make([]Word, n-len(tail)), tail...) }
	for i, c := range []struct {
		in, want []Word
	}{
		{[]Word{0, 1}, []Word{0, 1}},
		{pad(4, 0, 1), []Word{0, 1}},
		{pad(16, 0, 1), []Word{0, 1}},
		{pad(4, 1, 0, 0, 0), pad(4, 1, 0, 0, 0)},
		{pad(16, 1, 0, 0, 0), pad(4, 1, 0, 0, 0)},
		{pad(16, 1, 0, 0, 0, 0), pad(16, 1, 0, 0, 0, 0)},
		{pad(4), []Word{0, 0}},
		{pad(16), []Word{0, 0}},
	} {
		got, ok := IssuerNormal(c.in)
		if !ok || !slices.Equal(got, c.want) || !IssuerCanonical(got) {
			t.Fatalf("%d: %04X → %04X %v", i, c.in, got, ok)
		}
	}
	for _, bad := range [][]Word{nil, {1}, {0, 0, 1}, make([]Word, 8)} {
		if _, ok := IssuerNormal(bad); ok {
			t.Fatalf("%04X accepted", bad)
		}
	}
}
