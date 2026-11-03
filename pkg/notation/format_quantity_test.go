//ff:func feature=notation type=decoder control=sequence
//ff:what TestFormatQuantity — all four parts with a unit and bounds; a broken QUANTITY is false
package notation

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFormatQuantity(t *testing.T) {
	dec := func(c int64) core.Record { r, _ := core.EncodeDecimalValue(true, 0, c); return r }
	km := core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{0x000C, 0xA340}}
	q, _ := core.EncodeQuantity(core.Quantity{Amount: dec(652), Unit: km, Lower: dec(651), Upper: dec(653)})
	if s, ok := formatQuantity(q); !ok || s != "QUANTITY(652, Q828224, 651, 653)" {
		t.Fatal(s)
	}
	if _, ok := formatQuantity(core.Record{Kind: core.KindValue, Codec: core.CodecQuantity}); ok {
		t.Fatal("broken: want false")
	}
}
