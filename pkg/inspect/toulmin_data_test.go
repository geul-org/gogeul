//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestToulminData — SOURCE, TEXT, control, NODE and LINK decode; bad control and unknown types fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

func TestToulminData(t *testing.T) {
	w := toulmin.Wire{Span: 1}
	src, _ := toulmin.EncodeSource(toulmin.TextPacket{TID: 1, Bytes: []byte("s")}, toulmin.Wire{Span: 2})
	txt, _ := toulmin.EncodeText(toulmin.TextPacket{TID: 2, Bytes: []byte("t")}, w)
	sec, _ := toulmin.EncodeControl(toulmin.ControlPacket{ID: 3, Code: toulmin.ControlSection, Arg: 5}, w)
	node, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 4, Type: toulmin.NodeCase}, w)
	link, _ := toulmin.EncodeEdge(toulmin.EdgePacket{ID: 5, Prop: toulmin.PropRun, Src: 4, DstTID: 4}, w)
	for _, ws := range [][]core.Word{txt, sec, node, link, src} {
		r, _, _ := core.DecodeRecord(ws)
		ref, _ := toulmin.RecordRef(r)
		data, enc, err := toulminData(r, ref, &w)
		if err != nil || data == nil || len(enc) != len(ws) {
			t.Fatalf("%s: %+v %v", toulmin.RecordName(ref), data, err)
		}
	}
	if w.Span != 2 {
		t.Fatalf("SOURCE must set the span width: %d", w.Span)
	}
	bad := core.Record{Kind: core.KindEdge, Codec: 0x7C}
	if _, _, err := toulminData(bad, toulmin.CodecRef{Codec: 0x7C}, &w); err == nil {
		t.Fatal("bad control: want error")
	}
	if _, _, err := toulminData(core.Record{}, toulmin.CodecRef{Codec: toulmin.CodecPredicateDeclare}, &w); err == nil {
		t.Fatal("PREDICATE-DECLARE: want error")
	}
}
