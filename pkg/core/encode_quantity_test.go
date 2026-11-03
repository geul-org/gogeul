//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestEncodeQuantity — 652 km ±1 and a unitless 42 round-trip with their sizes; a bad record fails
package core

import (
	"math/big"
	"testing"
)

func TestEncodeQuantity(t *testing.T) {
	absent := Record{Kind: KindReference}
	km := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{0x000C, 0xA250}}
	dec := func(c int64) Record { return EncodeSignedDecimalBig(0, big.NewInt(c)) }
	for want, q := range map[int]Quantity{
		17: {Amount: dec(652), Unit: km, Lower: dec(651), Upper: dec(653)},
		8:  {Amount: dec(42), Unit: absent, Lower: absent, Upper: absent},
	} {
		r, err := EncodeQuantity(q)
		ws, _ := EncodeRecord(r)
		back, derr := DecodeQuantity(r)
		if err != nil || derr != nil || len(ws) != want || back.Unit.Codec != q.Unit.Codec || IsAbsent(back.Lower) != IsAbsent(q.Lower) {
			t.Fatalf("%d words: %d %v %v", want, len(ws), err, derr)
		}
	}
	if _, err := EncodeQuantity(Quantity{Amount: Record{Kind: KindValue, Codec: 0x80}}); err == nil {
		t.Fatal("bad codec: want error")
	}
}
