//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestZeroExpDecimal — SIGNED-/UNSIGNED-DECIMAL with E 0 are true; a non-zero exponent, an integer and a short payload are false
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestZeroExpDecimal(t *testing.T) {
	v := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindValue, Codec: c, Payload: p}
	}
	for _, r := range []core.Record{v(core.CodecSignedDecimal, 0, 0, 5), v(core.CodecUnsignedDecimal, 0, 0, 5)} {
		if !zeroExpDecimal(r) {
			t.Fatalf("%+v: want true", r)
		}
	}
	for _, r := range []core.Record{v(core.CodecSignedDecimal, 0xFFFF, 0xFFFF, 5), v(core.CodecSignedDecimal, 0, 1, 5), v(core.CodecSignedInteger, 0, 0), v(core.CodecUnsignedDecimal, 0)} {
		if zeroExpDecimal(r) {
			t.Fatalf("%+v: want false", r)
		}
	}
}
