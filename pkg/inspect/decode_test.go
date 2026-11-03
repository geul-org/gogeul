//ff:func feature=inspect type=decoder control=sequence
//ff:what TestDecode — a full document interprets HEAD and BODY; a HEAD SOURCE carries its Edge Id; odd bytes and a non-document fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestDecode(t *testing.T) {
	d, err := Decode(sampleDoc(t))
	if err != nil {
		t.Fatal(err)
	}
	if len(d.Head) != 6 || len(d.Body) != 8 || d.Profile != "tangeul-1.0" || !d.Toulmin {
		t.Fatalf("got %+v", d)
	}
	if _, ok := d.Body[3].Data.(toulmin.NodePacket); !ok || d.Body[3].Name != "CASE-DECLARE" || d.Body[4].ID != 12 {
		t.Fatalf("body: %+v", d.Body[3:5])
	}
	src := core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: []core.Word{0x0051, 3, 0x0051, 0, 0x0141, 42}}
	tr := core.Record{Kind: core.KindEdge, Codec: core.EdgeTriple, Payload: []core.Word{0x0051, 4}}
	b, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(core.SpecVersion), src}, []core.Record{tr})
	h, err := Decode(b)
	if err != nil || len(h.Head) != 2 || h.Head[1].ID != 3 || h.Head[1].Name != "SOURCE" || h.Body[0].ID != 4 || MaxLocalID(h) != 4 {
		t.Fatalf("HEAD SOURCE: %+v %v", h, err)
	}
	if _, err := Decode([]byte{1}); err == nil {
		t.Fatal("odd bytes: want error")
	}
	if _, err := Decode([]byte{0x1C, 0x40}); err == nil {
		t.Fatal("not a document: want error")
	}
}
