//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestPlaceDocRecord — HEAD/BODY/END routing, HEAD SOURCE and QUALIFIER, and every placement error
package core

import "testing"

func TestPlaceDocRecord(t *testing.T) {
	st := &docState{}
	meta := Record{Kind: KindMeta, Codec: MetaProfile}
	edge := Record{Kind: KindEdge, Codec: 0x60}
	for _, r := range []Record{meta, {Kind: KindMeta, Codec: MetaBody}, edge, {Kind: KindMeta, Codec: MetaEnd}} {
		if err := placeDocRecord(st, r); err != nil {
			t.Fatal(err)
		}
	}
	if len(st.doc.Head) != 1 || len(st.doc.Body) != 1 || !st.ended {
		t.Fatalf("routing: %+v", st)
	}
	if err := placeDocRecord(st, edge); err == nil {
		t.Fatal("record after END must fail")
	}
	implicit := &docState{}
	if err := placeDocRecord(implicit, edge); err != nil || !implicit.inBody {
		t.Fatal("non-Meta first record starts an implicit BODY")
	}
	if err := placeDocRecord(implicit, Record{Kind: KindMeta, Codec: MetaBody}); err == nil {
		t.Fatal("BODY after body records must fail")
	}
	if err := placeDocRecord(implicit, meta); err == nil {
		t.Fatal("Meta in BODY must fail")
	}
	if err := placeDocRecord(&docState{}, Record{Kind: KindMeta, Codec: MetaDoc1Word}); err == nil {
		t.Fatal("nested DOC must fail")
	}
	head := &docState{}
	_ = placeDocRecord(head, meta)
	if err := placeDocRecord(head, edge); err == nil {
		t.Fatal("non-Meta in HEAD must fail")
	}
	q42 := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{42}}
	p813 := Record{Kind: KindReference, Codec: CodecWikidataPID, Payload: []Word{813}}
	src := headEdgeFixture(EdgeSource, 1, EncodeLocalRef(0), q42)
	qual := headEdgeFixture(EdgeQualifier, 2, p813, EncodeLocalRef(1), q42)
	withSource := &docState{}
	for _, r := range []Record{src, meta, qual} {
		if err := placeDocRecord(withSource, r); err != nil {
			t.Fatal(err)
		}
	}
	if len(withSource.doc.Head) != 3 || withSource.inBody {
		t.Fatalf("HEAD SOURCE and its QUALIFIER, even first: %+v", withSource)
	}
	if err := placeDocRecord(withSource, headEdgeFixture(EdgeSource, 3, EncodeLocalRef(1), q42)); err == nil {
		t.Fatal("a SOURCE on a record in HEAD must fail (missing BODY)")
	}
}
