//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeQuantity — wrong kinds, leftovers, missing records and wrong record forms fail
package core

import (
	"math/big"
	"testing"
)

func TestDecodeQuantity(t *testing.T) {
	absent := Record{Kind: KindReference}
	dec := func(c int64) Record { return EncodeSignedDecimalBig(0, big.NewInt(c)) }
	good, _ := EncodeQuantity(Quantity{Amount: dec(1), Unit: absent, Lower: absent, Upper: absent})
	for name, r := range map[string]Record{
		"kind":     {Kind: KindValue, Codec: CodecCoordinateEarth},
		"short":    {Kind: KindValue, Codec: CodecQuantity, Payload: good.Payload[:4]},
		"leftover": {Kind: KindValue, Codec: CodecQuantity, Payload: append(append([]Word{}, good.Payload...), 0)},
		"amount":   {Kind: KindValue, Codec: CodecQuantity, Payload: append([]Word{0x0000}, good.Payload[4:]...)},
		"unit":     {Kind: KindValue, Codec: CodecQuantity, Payload: append(append(append([]Word{}, good.Payload[:4]...), 0x0AC0, 0x6162), good.Payload[5:]...)},
		"bound":    {Kind: KindValue, Codec: CodecQuantity, Payload: append(append(append([]Word{}, good.Payload[:5]...), 0x0AC0, 0x6162), good.Payload[6:]...)},
	} {
		if _, err := DecodeQuantity(r); err == nil {
			t.Fatalf("%s: want error", name)
		}
	}
}
