//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestNumberProblem — canonical integers, decimals, finite floats and valid fractions are ""; non-minimal, short, NaN and broken or invalid fractions are reported; other codecs are left to the decoders
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNumberProblem(t *testing.T) {
	v := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindValue, Codec: c, Payload: p}
	}
	good := []core.Record{
		v(core.CodecSignedInteger, 0xFFFF), v(core.CodecUnsignedInteger, 0x8000), v(core.CodecSignedDecimal, 0xFFFF, 0xFFFE, 0x0929),
		v(core.CodecUnsignedDecimal, 0xFFFF, 0xFFFF, 0x3111), v(core.CodecBinaryFloat, 0x3C00),
		v(core.CodecFraction, 0x0611, 2, 0x0621, 3), v(core.CodecUTF8TextEven),
	}
	for _, r := range good {
		if p := numberProblem(r); p != "" {
			t.Fatalf("%+v: %s", r, p)
		}
	}
	bad := []core.Record{
		v(core.CodecSignedInteger, 0, 1), v(core.CodecUnsignedInteger, 0, 1), v(core.CodecSignedDecimal, 0, 0),
		v(core.CodecSignedDecimal, 0, 0, 0xFFFF, 0xFFFF), v(core.CodecUnsignedDecimal, 0, 0, 0, 1),
		v(core.CodecBinaryFloat, 0x7E00), v(core.CodecFraction, 0x0621, 2, 0x0621, 3), v(core.CodecFraction, 0x0611, 2, 0x0621, 0),
	}
	for _, r := range bad {
		if numberProblem(r) == "" {
			t.Fatalf("%+v: want a problem", r)
		}
	}
}
