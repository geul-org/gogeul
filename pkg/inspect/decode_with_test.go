//ff:func feature=inspect type=decoder control=sequence
//ff:what TestDecodeWith — given specs lay out records, the document's own DECLAREs win; bad bytes fail
package inspect

import (
	"github.com/geul-org/gogeul/pkg/core"
	"slices"
	"testing"
)

func TestDecodeWith(t *testing.T) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	tomb := core.EncodePublication(core.Publication{Codec: core.EdgeTombstone, EdgeID: 3, Issuer: []core.Word{0, 0}, Doc: []core.Word{0, 0, 0, 1}, Version: 1, Reason: 1})
	in := slices.Concat([]core.Word{0x1C30}, enc(tomb), []core.Word{0x1C40})
	d, err := DecodeWith(core.PutWords(core.WrapDocument(in)), tombstoneSpec())
	if err != nil || len(Layout(d.Body[0])) != 4 || d.Body[0].ID != 3 {
		t.Fatalf("given: %+v %v", d.Body, err)
	}
	own := slices.Concat([]core.Word{0x1C30}, enc(tomb),
		enc(core.EncodeDeclare(core.Declare{EdgeID: 5, Target: core.TargetWord(core.KindEdge, core.EdgeTombstone)})),
		enc(core.EncodeFieldDecl(core.FieldDecl{EdgeID: 4, Declare: 5, Width: 7})),
		[]core.Word{0x1C40})
	d, _ = DecodeWith(core.PutWords(core.WrapDocument(own)), tombstoneSpec())
	if len(d.Body[0].Spec) != 1 {
		t.Fatalf("own spec: %+v", d.Body[0].Spec)
	}
	if _, err := DecodeWith([]byte{1}, nil); err == nil {
		t.Fatal("odd bytes: want error")
	}
}
