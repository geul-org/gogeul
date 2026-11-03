//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestFractionProblem — 2/3, 2/4 (not reduced) and 23.45/1256.1 are valid; non-minimal components, an exponent-0 DECIMAL and a zero denominator (0 or 0.0) are not
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFractionProblem(t *testing.T) {
	v := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindValue, Codec: c, Payload: p}
	}
	si, ui := v(core.CodecSignedInteger, 2), v(core.CodecUnsignedInteger, 3)
	sd, ud := v(core.CodecSignedDecimal, 0xFFFF, 0xFFFE, 0x0929), v(core.CodecUnsignedDecimal, 0xFFFF, 0xFFFF, 0x3111)
	for _, f := range []core.Fraction{{Numerator: si, Denominator: ui}, {Numerator: si, Denominator: v(core.CodecUnsignedInteger, 4)}, {Numerator: sd, Denominator: ud}} {
		if p := fractionProblem(f); p != "" {
			t.Fatalf("%+v: %s", f, p)
		}
	}
	bad := []core.Fraction{
		{Numerator: v(core.CodecSignedInteger, 0, 2), Denominator: ui},
		{Numerator: si, Denominator: v(core.CodecUnsignedInteger, 0, 3)},
		{Numerator: v(core.CodecSignedDecimal, 0, 0, 5), Denominator: ui},
		{Numerator: si, Denominator: v(core.CodecUnsignedDecimal, 0, 0, 5)},
		{Numerator: si, Denominator: v(core.CodecUnsignedInteger, 0)},
		{Numerator: si, Denominator: v(core.CodecUnsignedDecimal, 0xFFFF, 0xFFFF, 0)},
		{Numerator: si, Denominator: v(core.CodecBinaryFloat, 0x3C00)},
	}
	for i, f := range bad {
		if fractionProblem(f) == "" {
			t.Fatalf("case %d: want a problem", i)
		}
	}
}
