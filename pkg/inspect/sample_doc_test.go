//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what sampleDoc — test fixture: a TANGEUL-profile document with every HEAD form and BODY record kind
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// sampleDoc builds a document whose HEAD has every typed Meta plus an unknown
// one, and whose BODY has a SOURCE, TEXT, framing, a NODE and a LINK, a
// PREDICATE-DECLARE (raw), a standalone Value and an unknown Value.
func sampleDoc(t *testing.T) []byte {
	t.Helper()
	w := toulmin.Wire{Span: 1}
	meta := func(c uint16, v core.Record) []core.Word {
		vw, _ := core.EncodeRecord(v)
		ws, _ := core.EncodeRecord(core.Record{Kind: core.KindMeta, Codec: c, Payload: vw})
		return ws
	}
	raw := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	var in []core.Word
	in = append(in, raw(core.EncodeSpecVersion(core.SpecVersion))...)
	in = append(in, meta(core.MetaProfile, core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("tangeul-1.0")))...)
	in = append(in, meta(core.MetaDefaultLanguage, core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("ko")))...)
	in = append(in, raw(func() core.Record { r, _ := core.EncodeDocID([]core.Word{0, 2}, []core.Word{0, 0, 0, 5}); return r }())...)
	in = append(in, raw(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x0418, 0, 2, 0, 0, 0, 2, 0, 0, 1}})...)
	in = append(in, raw(core.Record{Kind: core.KindMeta, Codec: 0x2A, Payload: []core.Word{7}})...)
	in = append(in, 0x1C30)
	src, _ := toulmin.EncodeSource(toulmin.TextPacket{TID: 9, Bytes: []byte("# t\n")}, w)
	txt, _ := toulmin.EncodeText(toulmin.TextPacket{TID: 10, Bytes: []byte("n")}, w)
	sec, _ := toulmin.EncodeControl(toulmin.ControlPacket{ID: 11, Code: toulmin.ControlSection, Arg: 5}, w)
	node, _ := toulmin.EncodeNode(toulmin.NodePacket{TID: 1, Type: toulmin.NodeCase, Span: toulmin.Span{Off: 0, Len: 4}}, w)
	link, _ := toulmin.EncodeEdge(toulmin.EdgePacket{ID: 12, Prop: toulmin.PropLabel, Src: 1, DstTID: 10, Span: toulmin.Span{Len: 4}}, w)
	pred := raw(core.Record{Kind: core.KindEdge, Codec: toulmin.CodecPredicateDeclare, Payload: []core.Word{0x0051, 13, 0}})
	for _, p := range [][]core.Word{src, txt, sec, node, link, pred, raw(core.EncodeUintValue(5)), raw(core.Record{Kind: core.KindValue, Codec: 0x03, Payload: []core.Word{0x3c00}})} {
		in = append(in, p...)
	}
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
