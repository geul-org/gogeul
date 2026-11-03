//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestPropMovedRef — every relation moved from a GEUL format maps to its direct code 0x12BL~0x134L; ATTACKS does not
package toulmin

import "testing"

func TestPropMovedRef(t *testing.T) {
	want := map[PropCode]uint16{
		PropRequires: CodecRequires, PropMeans: CodecMeans, PropHas: CodecHas,
		PropRuns: CodecRuns, PropEvaluates: CodecEvaluates, PropTriggers: CodecTriggers,
		PropUntil: CodecUntil, PropSameAs: CodecSameAs, PropImportedFrom: CodecImportedFrom,
		PropLabel: CodecLabel,
	}
	for p, f := range want {
		if ref, err := propMovedRef(p); err != nil || ref.Codec != f {
			t.Fatalf("0x%04X: %+v %v", uint16(p), ref, err)
		}
	}
	if _, err := propMovedRef(PropAttacks); err == nil {
		t.Fatal("attacks is a TOULMIN 0x100L code, not a moved one")
	}
}
