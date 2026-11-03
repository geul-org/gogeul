//ff:func feature=resolve type=decoder control=sequence
//ff:what TestParseDocumentNames — names map to documents and codec TAGs to spec documents, Issuer BINDs and other records ignored; odd bytes and a non-document fail
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseDocumentNames(t *testing.T) {
	b := documentsDoc(map[string]core.Word{"sample-a": 1, "sample-b": 2})
	names, err := ParseDocumentNames(b)
	if err != nil || len(names) != 2 || names["sample-b"] != (DocKey{Issuer: "0x00000000", Doc: "0xffff000000000002"}) {
		t.Fatalf("got %v %v", names, err)
	}
	if n, _ := ParseDocumentNames(issuersDoc([2]any{core.Word(0), "geul.org"})); len(n) != 1 || n["geul.org"] != (DocKey{}) || n["a-document"].Doc != "0x0000000000000005" {
		t.Fatalf("issuer BIND skipped, document BIND kept: %v", n)
	}
	enc := func(r core.Record) []core.Word { ws, _ := core.EncodeRecord(r); return ws }
	p := append(enc(core.EncodeLocalRef(1)), 0x0980)
	p = append(p, enc(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("publication")))...)
	codecs := append([]core.Word{0x1C30}, enc(core.Record{Kind: core.KindEdge, Codec: core.EdgeBind, Payload: p})...)
	if n, _ := ParseDocumentNames(core.PutWords(core.WrapDocument(append(codecs, 0x1C40)))); n["publication"] != codecKey(0x0980) {
		t.Fatalf("codec: %v", n)
	}
	if _, err := ParseDocumentNames([]byte{1}); err == nil {
		t.Fatal("odd: want error")
	}
	if _, err := ParseDocumentNames([]byte{0x1C, 0x40}); err == nil {
		t.Fatal("not a document: want error")
	}
}
