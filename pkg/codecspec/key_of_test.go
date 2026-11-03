//ff:func feature=codecspec type=helper control=iteration dimension=1
//ff:what TestKeyOf — every key is the record's own code: Edges, and Reference 5Q codes with their Quantifier, pointers and REFERENCE-ISSUER alike
package codecspec

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestKeyOf(t *testing.T) {
	if k := keyOf(core.Record{Kind: core.KindEdge, Codec: core.EdgePublication}); k != 0x0B80 {
		t.Fatalf("edge: %04X", k)
	}
	ref := map[uint16]core.Word{
		2: 0x0020, 5: 0x0050, 6: 0x0060, 9: 0x0090, 0x1D: 0x01D0, 0x1E: 0x01E0,
		0x20: 0x0200, 0x24: 0x0240, 0x2E: 0x02E0, 0x3D: 0x03D0, 0x41: 0x0410, 0x5D: 0x05D0,
	}
	for c, want := range ref {
		if k := keyOf(core.Record{Kind: core.KindReference, Codec: c}); k != want {
			t.Fatalf("reference C 0x%02X: %04X want %04X", c, k, want)
		}
	}
}
