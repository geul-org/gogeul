//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestZoneCounts — assigned codes per zone match migration.md §1 (§3.14, §3.15 included): Reference 64 of 96 (the rest reserved R1), Value 37 (34 direct + 3 ISSUER), Edge 84 (81 direct + 3 ISSUER), Meta 42
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestZoneCounts(t *testing.T) {
	ref, value, edge, meta := 0, 0, 0, 0
	for c := uint16(0); c <= core.CodecMax(core.KindReference); c++ {
		r := recOf(core.Record{Kind: core.KindReference, Codec: c, Payload: []core.Word{1}})
		if core.IsSlotMarkerCodec(c) || rulesOf(checkReference(r, nil)) != "R1" {
			ref++
		}
	}
	for c := uint16(0); c <= 0x3F; c++ {
		if rulesOf(checkValue(recOf(core.Record{Kind: core.KindValue, Codec: c}))) != "V1" {
			value++
		}
		if metaDefined(c) {
			meta++
		}
	}
	for c := uint16(0); c <= core.CodecMax(core.KindEdge); c++ {
		if edgeDefined(c) {
			edge++
		}
	}
	if ref != 64 || value != 37 || edge != 84 || meta != 42 {
		t.Fatalf("Reference %d, Value %d, Edge %d, Meta %d; want 64, 37, 84, 42", ref, value, edge, meta)
	}
}
