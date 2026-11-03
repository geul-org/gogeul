//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeFraction — 2/3 decodes to its two records; wrong kind or codec, a missing record, leftovers and wrong component codecs fail
package core

import (
	"math/big"
	"testing"
)

func TestDecodeFraction(t *testing.T) {
	n, _ := EncodeIntegerBig(true, big.NewInt(2))
	d, _ := EncodeIntegerBig(false, big.NewInt(3))
	good, _ := EncodeFraction(Fraction{Numerator: n, Denominator: d})
	f, err := DecodeFraction(good)
	if err != nil || f.Numerator.Codec != CodecSignedInteger || f.Denominator.Codec != CodecUnsignedInteger || f.Denominator.Payload[0] != 3 {
		t.Fatalf("2/3: %+v %v", f, err)
	}
	for name, r := range map[string]Record{
		"kind":     {Kind: KindEdge, Codec: CodecFraction},
		"codec":    {Kind: KindValue, Codec: CodecQuantity},
		"short":    {Kind: KindValue, Codec: CodecFraction, Payload: good.Payload[:2]},
		"leftover": {Kind: KindValue, Codec: CodecFraction, Payload: append(append([]Word{}, good.Payload...), 0)},
		"swapped":  {Kind: KindValue, Codec: CodecFraction, Payload: []Word{0x0621, 3, 0x0611, 2}},
	} {
		if _, err := DecodeFraction(r); err == nil {
			t.Fatalf("%s: want error", name)
		}
	}
}
