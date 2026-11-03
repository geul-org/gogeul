//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what TestParseFractionPart — digits as INTEGER, a point or exponent as DECIMAL (signed or unsigned); a non-number, a bad decimal, exponent 0 and a negative denominator fail
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseFractionPart(t *testing.T) {
	cases := []struct {
		arg    string
		signed bool
		want   uint16
	}{{"-2", true, core.CodecSignedInteger}, {"3", false, core.CodecUnsignedInteger}, {"23.45", true, core.CodecSignedDecimal}, {"1256.1", false, core.CodecUnsignedDecimal}, {"12e2", false, core.CodecUnsignedDecimal}}
	for _, c := range cases {
		if r, err := parseFractionPart(c.arg, c.signed); err != nil || r.Codec != c.want {
			t.Fatalf("%s: %+v %v", c.arg, r, err)
		}
	}
	for _, s := range []string{"x", "1.2.3", "5.", "5e0", "-3"} {
		if _, err := parseFractionPart(s, false); err == nil {
			t.Fatalf("%s: want error", s)
		}
	}
}
