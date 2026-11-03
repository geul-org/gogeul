//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestQuantityNumbers — canonical amount and bounds (ABSENT skipped) are ""; a non-minimal amount, a NaN bound and an invalid FRACTION bound are reported
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestQuantityNumbers(t *testing.T) {
	v := func(c uint16, p ...core.Word) core.Record {
		return core.Record{Kind: core.KindValue, Codec: c, Payload: p}
	}
	absent := core.Record{Kind: core.KindReference}
	d := v(core.CodecSignedDecimal, 0, 0, 1)
	if p := quantityNumbers(core.Quantity{Amount: d, Unit: absent, Lower: v(core.CodecSignedInteger, 1), Upper: absent}); p != "" {
		t.Fatal(p)
	}
	for i, q := range []core.Quantity{
		{Amount: v(core.CodecSignedDecimal, 0, 0, 0, 1), Unit: absent, Lower: absent, Upper: absent},
		{Amount: d, Unit: absent, Lower: v(core.CodecBinaryFloat, 0x7FC0, 0), Upper: absent},
		{Amount: d, Unit: absent, Lower: absent, Upper: v(core.CodecFraction, 0x0411, 1, 0x0421, 0)},
	} {
		if quantityNumbers(q) == "" {
			t.Fatalf("case %d: want a problem", i)
		}
	}
}
