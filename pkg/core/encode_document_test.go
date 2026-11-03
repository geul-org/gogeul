//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDocument — HEAD and BODY records split back out; an empty document frames; an unencodable record fails
package core

import "testing"

func TestEncodeDocument(t *testing.T) {
	b, err := EncodeDocument([]Record{EncodeSpecVersion(2)}, []Record{EncodeLocalRef(1)})
	if err != nil {
		t.Fatal(err)
	}
	ws, _ := ReadWords(b)
	doc, err := SplitDocument(ws)
	if err != nil || len(doc.Head) != 1 || len(doc.Body) != 1 {
		t.Fatalf("%+v %v", doc, err)
	}
	if b, err := EncodeDocument(nil, nil); err != nil || len(b) == 0 {
		t.Fatal(err)
	}
	if _, err := EncodeDocument([]Record{{Kind: KindEdge, Codec: 0x140}}, nil); err == nil {
		t.Fatal("bad record: want error")
	}
}
