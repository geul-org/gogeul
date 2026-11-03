//ff:func feature=codecspec type=helper control=iteration dimension=1
//ff:what TestKeyOf — Edge keys keep their code; a quantified Reference folds to its scheme head only for 5 <= C < 0x3C (§7.3 boundaries)
package codecspec

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestKeyOf(t *testing.T) {
	if k := keyOf(core.Record{Kind: core.KindEdge, Codec: core.EdgePublication}); k != 0x0980 {
		t.Fatalf("edge: %04X", k)
	}
	ref := map[uint16]core.Word{
		2: 0x0020, 3: 0x0030, 4: 0x0040, 5: 0x0050, 9: 0x0050, 0x1D: 0x0190, 0x1E: 0x01E0, 0x2C: 0x0280, 0x2D: 0x02D0,
		0x36: 0x0320, 0x37: 0x0370, 0x3B: 0x0370, 0x3C: 0x03C0, 0x3D: 0x03D0, 0x3E: 0x03E0, 0x3F: 0x03F0,
	}
	for c, want := range ref {
		if k := keyOf(core.Record{Kind: core.KindReference, Codec: c}); k != want {
			t.Fatalf("reference C 0x%02X: %04X want %04X", c, k, want)
		}
	}
}
