//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestQuantity — 652 Qn [651, 653], a unitless 1.20 and integer/FRACTION/BINARY-FLOAT parts read as text; a bad QUANTITY and a non-minimal bound fail
package inspect

import (
	"math/big"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestQuantity(t *testing.T) {
	absent := core.Record{Kind: core.KindReference}
	km := core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{0x000C, 0xA250}}
	dec := func(e int32, c int64) core.Record { return core.EncodeSignedDecimalBig(e, big.NewInt(c)) }
	r, _ := core.EncodeQuantity(core.Quantity{Amount: dec(0, 652), Unit: km, Lower: dec(0, 651), Upper: dec(0, 653)})
	if a, u, l, h, err := quantity(r); err != nil || a != "652" || u != "Q827984" || l != "651" || h != "653" {
		t.Fatalf("km: %q %q %q %q %v", a, u, l, h, err)
	}
	r, _ = core.EncodeQuantity(core.Quantity{Amount: dec(-2, 120), Unit: absent, Lower: absent, Upper: absent})
	if a, u, l, h, err := quantity(r); err != nil || a != "1.20" || u != "" || l != "" || h != "" {
		t.Fatalf("unitless: %q %q %q %q %v", a, u, l, h, err)
	}
	frac := core.Record{Kind: core.KindValue, Codec: core.CodecFraction, Payload: []core.Word{0x0611, 2, 0x0621, 3}}
	one := core.Record{Kind: core.KindValue, Codec: core.CodecSignedInteger, Payload: []core.Word{1}}
	bf := core.Record{Kind: core.KindValue, Codec: core.CodecBinaryFloat, Payload: []core.Word{0x3C00}}
	r, _ = core.EncodeQuantity(core.Quantity{Amount: frac, Unit: absent, Lower: one, Upper: bf})
	if a, _, l, h, err := quantity(r); err != nil || a != "2/3" || l != "1" || h != "1" {
		t.Fatalf("numbers: %q %q %q %v", a, l, h, err)
	}
	if _, _, _, _, err := quantity(core.Record{Kind: core.KindValue, Codec: core.CodecQuantity}); err == nil {
		t.Fatal("empty: want error")
	}
	loose := core.Record{Kind: core.KindValue, Codec: core.CodecSignedDecimal, Payload: []core.Word{0, 0, 0, 1}}
	for _, q := range []core.Quantity{{Amount: loose, Unit: absent, Lower: absent, Upper: absent}, {Amount: dec(0, 1), Unit: absent, Lower: loose, Upper: absent}, {Amount: dec(0, 1), Unit: absent, Lower: absent, Upper: loose}} {
		r, _ := core.EncodeQuantity(q)
		if _, _, _, _, err := quantity(r); err == nil {
			t.Fatal("non-minimal: want error")
		}
	}
}
