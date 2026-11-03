//ff:func feature=page type=builder control=sequence
//ff:what TestBodyRow — Edge Id anchor, label, back-references, span and words; position anchor without an Edge Id
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestBodyRow(t *testing.T) {
	e := inspect.Entry{Name: "ATTACKS", ID: 7, Words: []core.Word{0x1268, 7}, Record: core.Record{Payload: []core.Word{7}},
		Data: toulmin.EdgePacket{ID: 7, Src: 1, DstTID: 2, Span: toulmin.Span{Off: 3, Len: 4}}}
	r := bodyRow(5, e, map[uint64]string{7: "공격"}, []uint64{9}, 2)
	if r.Anchor != "e7" || r.Label != "공격" || r.Class != "link" || len(r.Back) != 1 || !r.HasSpan || r.SpanLen != 4 || r.HeadParts[0].Lines[0].Word != "1268" || r.BodyParts[0].Lines[0].Word != "0007" || r.Depth != 2 || r.Index != 5 {
		t.Fatalf("got %+v", r)
	}
	if r := bodyRow(3, inspect.Entry{Words: []core.Word{0x0A60}, Record: core.Record{Kind: core.KindValue, Codec: 0x60}}, nil, nil, 0); r.Anchor != "r3" || r.HasSpan || r.Name != "VALUE 0x60" {
		t.Fatalf("bare: %+v", r)
	}
}
