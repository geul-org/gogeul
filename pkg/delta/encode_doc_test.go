//ff:func feature=delta type=encoder control=sequence
//ff:what TestEncodeDoc — HEAD and BODY round-trip through core.SplitDocument; an empty HEAD writes no BODY marker
package delta

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEncodeDoc(t *testing.T) {
	ws, _ := core.ReadWords(encodeDoc([]core.Record{core.EncodeSpecVersion(2)}, []core.Record{bindOf(1, 2, "a")}))
	d, err := core.SplitDocument(ws)
	if err != nil || len(d.Head) != 1 || len(d.Body) != 1 || d.Body[0].Codec != core.EdgeBind {
		t.Fatalf("%+v %v", d, err)
	}
	ws, _ = core.ReadWords(encodeDoc(nil, []core.Record{bindOf(1, 2, "a")}))
	if d, err := core.SplitDocument(ws); err != nil || len(d.Body) != 1 || slices.Contains(ws, core.TargetWord(core.KindMeta, core.MetaBody)) {
		t.Fatalf("empty HEAD: %04X %v", ws, err)
	}
}
