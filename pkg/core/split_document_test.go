//ff:func feature=core type=decoder control=sequence
//ff:what TestSplitDocument — whole-document parse plus bad open, bad record, placement and missing END errors
package core

import "testing"

func TestSplitDocument(t *testing.T) {
	inner := []Word{0x1C60, 0x1C30}
	edge, _ := EncodeRecord(Record{Kind: KindEdge, Codec: 0x60, Payload: []Word{1}})
	inner = append(append(inner, edge...), 0x1C40)
	ws := WrapDocument(inner)
	d, err := SplitDocument(ws)
	if err != nil || len(d.Head) != 1 || len(d.Body) != 1 {
		t.Fatalf("got %+v %v", d, err)
	}
	if _, err := SplitDocument(nil); err == nil {
		t.Fatal("empty input must fail")
	}
	if _, err := SplitDocument([]Word{0x1C40}); err == nil {
		t.Fatal("non-DOC open must fail")
	}
	noEnd := WrapDocument([]Word{0x1C30})
	if _, err := SplitDocument(noEnd); err == nil {
		t.Fatal("missing END must fail")
	}
	badRec := WrapDocument([]Word{0xE000})
	if _, err := SplitDocument(badRec); err == nil {
		t.Fatal("bad record must fail")
	}
	twice := WrapDocument([]Word{0x1C40, 0x1C40})
	if _, err := SplitDocument(twice); err == nil {
		t.Fatal("record after END must fail")
	}
}
