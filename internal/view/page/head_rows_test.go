//ff:func feature=page type=builder control=sequence
//ff:what TestHeadRows — each Meta is a `meta` row anchored h<n> with its fields and words; a HEAD SOURCE is an Edge row anchored by its Edge Id
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestHeadRows(t *testing.T) {
	rows := headRows(fixtureDoc(t).Head)
	if len(rows) != 2 || rows[1].Anchor != "h1" || rows[1].Name != "PROFILE" || rows[1].Class != "meta" || rows[1].Fields[0].Text != `"tangeul-1.0"` || len(rows[1].HeadParts) != 1 {
		t.Fatalf("got %+v", rows)
	}
	b, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(core.SpecVersion), {Kind: core.KindEdge, Codec: core.EdgeSource, Payload: []core.Word{0x0051, 3, 0x0051, 0, 0x0141, 42}}}, nil)
	d, err := inspect.Decode(b)
	if err != nil {
		t.Fatal(err)
	}
	rows = headRows(d.Head)
	if len(rows) != 2 || rows[1].Anchor != "e3" || rows[1].ID != 3 || rows[1].Name != "SOURCE" || !rows[1].Edge {
		t.Fatalf("HEAD SOURCE: %+v", rows[1])
	}
}
