//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestRetiredInHeadSource — a HEAD SOURCE added in the next version (Target #0, the document itself) takes its Edge Id from NEXT-LOCAL-ID and passes; #0 is never a retired ID
package resolve

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestRetiredInHeadSource(t *testing.T) {
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	doc, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7})
	src := slices.Concat(enc(core.EncodeLocalRef(3)), enc(core.EncodeLocalRef(0)), enc(core.Record{Kind: core.KindReference, Codec: core.CodecWikidataQID, Payload: []core.Word{2013}}))
	in := slices.Concat(enc(doc), enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: src}), []core.Word{0x1C30})
	for i, txt := range []string{"a", "b"} {
		p := slices.Concat(enc(core.EncodeLocalRef(uint64(i+1))), enc(core.EncodeLocalRef(uint64(i+1))), enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(txt))))
		in = append(in, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p})...)
	}
	next := core.PutWords(core.WrapDocument(append(in, 0x1C40)))
	body := edgeDoc(7, "a", "b")
	nd, err := inspect.Decode(next)
	if err != nil || len(nd.Head) == 0 || !slices.Contains(inspect.LocalIDs(nd), 0) {
		t.Fatal("fixture: the HEAD SOURCE should mention #0")
	}
	if id, ok := retiredIn(body, next, 3); ok {
		t.Fatalf("HEAD SOURCE #3 on #0: got retired #%d", id)
	}
	if id, ok := retiredIn(edgeDoc(7, "a"), next, 3); !ok || id != 2 {
		t.Fatalf("#2 dropped is still found: %d %v", id, ok)
	}
}
