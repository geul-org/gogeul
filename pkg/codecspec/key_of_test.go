//ff:func feature=codecspec type=helper control=iteration dimension=1
//ff:what TestKeyOf — Edge keys keep their code; a quantified Reference folds to its scheme head only for the 5Q codes of sections 0 and 1 (§7.3 boundaries)
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
		2: 0x0020, 3: 0x0030, 4: 0x0040, 5: 0x0050, 9: 0x0050, 0x1D: 0x0190, 0x1E: 0x01E0, 0x1F: 0x01F0,
		0x20: 0x0200, 0x24: 0x0200, 0x2E: 0x02A0, 0x2F: 0x02F0, 0x3D: 0x0390, 0x3E: 0x03E0, 0x3F: 0x03F0,
		0x40: 0x0400, 0x41: 0x0410, 0x42: 0x0420, 0x43: 0x0430, 0x5C: 0x05C0, 0x5D: 0x05D0, 0x5F: 0x05F0,
	}
	for c, want := range ref {
		if k := keyOf(core.Record{Kind: core.KindReference, Codec: c}); k != want {
			t.Fatalf("reference C 0x%02X: %04X want %04X", c, k, want)
		}
	}
}
