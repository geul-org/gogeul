//ff:func feature=page type=helper control=iteration dimension=1
//ff:what TestTagName — slot markers, 5Q Reference with Quantifier, spec names, the fallback and hex
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTagName(t *testing.T) {
	cases := []struct {
		k        core.Kind
		c        uint16
		fallback string
		want     string
	}{
		{core.KindReference, 1, "", "자리 표지 UNKNOWN-VALUE"},
		{core.KindReference, 0x05, "", "LOCAL · " + core.QuantifierName(0x05)},
		{core.KindValue, core.CodecBitVector, "", "BIT-VECTOR"},
		{core.KindEdge, 0x89, "ALIAS", "ALIAS"},
		{core.KindEdge, 0x95, "", "0x095"},
	}
	for _, c := range cases {
		if got := tagName(c.k, c.c, c.fallback); got != c.want {
			t.Fatalf("%v %X: %q, want %q", c.k, c.c, got, c.want)
		}
	}
}
