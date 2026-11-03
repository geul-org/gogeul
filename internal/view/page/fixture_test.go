//ff:func feature=page type=builder control=iteration dimension=1
//ff:what fixtureDoc — test fixture: a TANGEUL document with SOURCE, labelled case, rule, links, a case block and raw records
package page

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// fixtureDoc builds and interprets a small TANGEUL document: SOURCE, a SECTION,
// a CASE-DECLARE #1 labelled "사안" by TEXT #3, a RULE #2 that checks #1,
// CASE-BEGIN/CASE-END around an ATTACKS link, and a PREDICATE-DECLARE that
// stays raw.
func fixtureDoc(t *testing.T) inspect.Document {
	t.Helper()
	w := toulmin.Wire{Span: 1}
	meta := func(c uint16, v core.Record) []core.Word {
		vw, _ := core.EncodeRecord(v)
		ws, _ := core.EncodeRecord(core.Record{Kind: core.KindMeta, Codec: c, Payload: vw})
		return ws
	}
	var in []core.Word
	sv, _ := core.EncodeRecord(core.EncodeSpecVersion(toulmin.SpecVersion))
	in = append(in, sv...)
	in = append(in, meta(core.MetaProfile, core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(toulmin.ProfileString())))...)
	in = append(in, 0x1C30)
	check := toulmin.TID(1)
	src, _ := toulmin.EncodeSource(toulmin.TextPacket{TID: 10, Bytes: []byte("# 사안\n규칙")}, w)
	sec, _ := toulmin.EncodeControl(toulmin.ControlPacket{ID: 11, Code: toulmin.ControlSection, Arg: 5}, w)
	cs, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 1, Type: toulmin.NodeCase, Span: toulmin.Span{Off: 2, Len: 6}}, w)
	txt, _ := toulmin.EncodeText(toulmin.TextPacket{TID: 3, Bytes: []byte("사안")}, w)
	lbl, _ := toulmin.EncodeEdge(toulmin.EdgePacket{ID: 12, Prop: toulmin.PropLabel, Src: 1, DstTID: 3}, w)
	rule, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 2, Type: toulmin.NodeRule, Flags: toulmin.NodeFlagCheck, Check: &check, Span: toulmin.Span{Off: 9, Len: 6}}, w)
	begin, _ := toulmin.EncodeControl(toulmin.ControlPacket{ID: 13, Code: toulmin.ControlCaseBegin, Arg: 1}, w)
	att, _ := toulmin.EncodeEdge(toulmin.EdgePacket{ID: 14, Prop: toulmin.PropAttacks, Src: 2, DstTID: 1}, w)
	end, _ := toulmin.EncodeControl(toulmin.ControlPacket{ID: 15, Code: toulmin.ControlCaseEnd, Arg: 1}, w)
	pred, _ := core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: toulmin.CodecPredicateDeclare, Payload: []core.Word{0x0051, 16, 0}})
	for _, p := range [][]core.Word{src, sec, cs, txt, lbl, rule, begin, att, end, pred} {
		in = append(in, p...)
	}
	d, err := inspect.Decode(core.PutWords(core.WrapDocument(append(in, 0x1C40))))
	if err != nil {
		t.Fatal(err)
	}
	return d
}
