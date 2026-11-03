//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestZoneCounts — assigned codes per zone match migration.md §1 (§3.14 included): Value 37 (34 direct + 3 ISSUER), Edge 84 (81 direct + 3 ISSUER), Meta 42
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestZoneCounts(t *testing.T) {
	value, edge, meta := 0, 0, 0
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
	if value != 37 || edge != 84 || meta != 42 {
		t.Fatalf("Value %d, Edge %d, Meta %d; want 37, 84, 42", value, edge, meta)
	}
}
