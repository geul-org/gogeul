//ff:func feature=notation type=decoder control=iteration dimension=1
//ff:what TestFractionPartText — INTEGER as digits, DECIMAL with a point or exponent; an exponent-0 DECIMAL and broken components are false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFractionPartText(t *testing.T) {
	v := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindValue, Codec: c, Payload: p}
	}
	cases := map[string]core.Record{
		"-2":     v(core.CodecSignedInteger, 0xFFFE),
		"3":      v(core.CodecUnsignedInteger, 3),
		"23.45":  v(core.CodecSignedDecimal, 0xFFFF, 0xFFFE, 0x0929),
		"1256.1": v(core.CodecUnsignedDecimal, 0xFFFF, 0xFFFF, 0x3111),
		"12e2":   v(core.CodecSignedDecimal, 0, 2, 12),
	}
	for want, r := range cases {
		if got, ok := fractionPartText(r); !ok || got != want {
			t.Fatalf("%s: got %s %v", want, got, ok)
		}
	}
	for _, r := range []core.Record{v(core.CodecSignedDecimal, 0, 0, 5), v(core.CodecUnsignedInteger, 0, 1), v(core.CodecSignedDecimal, 0)} {
		if s, ok := fractionPartText(r); ok {
			t.Fatalf("%+v: want false, got %s", r, s)
		}
	}
}
