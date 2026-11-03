//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestNumberText — integers past 64 bits, decimals with their quantum, binary16 and FRACTION as written; broken payloads, binary128 and non-number codecs fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNumberText(t *testing.T) {
	v := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindValue, Codec: c, Payload: p}
	}
	cases := map[string]core.Record{
		"18446744073709551616": v(core.CodecUnsignedInteger, 1, 0, 0, 0, 0),
		"-1":                   v(core.CodecSignedInteger, 0xFFFF),
		"1.20":                 v(core.CodecSignedDecimal, 0xFFFF, 0xFFFE, 120),
		"1256.1":               v(core.CodecUnsignedDecimal, 0xFFFF, 0xFFFF, 0x3111),
		"1":                    v(core.CodecBinaryFloat, 0x3C00),
		"23.45/1256.1":         v(core.CodecFraction, 0x0483, 0xFFFF, 0xFFFE, 0x0929, 0x0493, 0xFFFF, 0xFFFF, 0x3111),
	}
	for want, r := range cases {
		if got, err := numberText(r); err != nil || got != want {
			t.Fatalf("%s: got %q %v", want, got, err)
		}
	}
	for _, r := range []core.Record{v(core.CodecSignedInteger, 0, 1), v(core.CodecSignedDecimal, 0), v(core.CodecBinaryFloat, make([]core.Word, 8)...), v(core.CodecUTF8TextEven)} {
		if _, err := numberText(r); err == nil {
			t.Fatalf("%+v: want error", r)
		}
	}
}
