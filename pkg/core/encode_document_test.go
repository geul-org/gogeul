//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDocument — HEAD and BODY records split back out; an empty HEAD omits BODY (DOC END for an empty document); an unencodable record fails
package core

import (
	"slices"
	"testing"
)

func TestEncodeDocument(t *testing.T) {
	b, err := EncodeDocument([]Record{EncodeSpecVersion(2)}, []Record{EncodeLocalRef(1)})
	if err != nil {
		t.Fatal(err)
	}
	ws, _ := ReadWords(b)
	doc, err := SplitDocument(ws)
	if err != nil || len(doc.Head) != 1 || len(doc.Body) != 1 || !slices.Contains(ws, TargetWord(KindMeta, MetaBody)) {
		t.Fatalf("%+v %v", doc, err)
	}
	b, err = EncodeDocument(nil, []Record{EncodeLocalRef(1)})
	ws, _ = ReadWords(b)
	if err != nil || slices.Contains(ws, TargetWord(KindMeta, MetaBody)) {
		t.Fatalf("empty HEAD keeps BODY: %04X %v", ws, err)
	}
	if doc, err := SplitDocument(ws); err != nil || len(doc.Head) != 0 || len(doc.Body) != 1 {
		t.Fatalf("%+v %v", doc, err)
	}
	b, err = EncodeDocument(nil, nil)
	ws, _ = ReadWords(b)
	if err != nil || !slices.Equal(ws, []Word{0x1C01, 3, TargetWord(KindMeta, MetaEnd)}) {
		t.Fatalf("empty document: %04X %v", ws, err)
	}
	if _, err := EncodeDocument([]Record{{Kind: KindEdge, Codec: 0x140}}, nil); err == nil {
		t.Fatal("bad record: want error")
	}
}
