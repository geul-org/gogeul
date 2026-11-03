//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestPropFormatRef — every format-number relation maps to its format; a direct-code relation does not
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestPropFormatRef(t *testing.T) {
	want := map[PropCode]uint32{
		PropRequires: FormatRequires, PropMeans: FormatMeans, PropHas: FormatHas,
		PropRuns: FormatRuns, PropEvaluates: FormatEvaluates, PropTriggers: FormatTriggers,
		PropUntil: FormatUntil, PropSameAs: FormatSameAs, PropImportedFrom: FormatImportedFrom,
		PropLabel: FormatLabel,
	}
	for p, f := range want {
		if ref, err := propFormatRef(p); err != nil || ref.Codec != core.CodecEdgeIssuer2Word || ref.Format != f {
			t.Fatalf("0x%04X: %+v %v", uint16(p), ref, err)
		}
	}
	if _, err := propFormatRef(PropAttacks); err == nil {
		t.Fatal("attacks has a direct code, not a format")
	}
}
