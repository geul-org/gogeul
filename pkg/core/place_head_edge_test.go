//ff:func feature=core type=decoder control=sequence
//ff:what TestPlaceHeadEdge — a document SOURCE goes to the HEAD and lets its QUALIFIER follow; a QUALIFIER on a QUALIFIER and other records do not
package core

import "testing"

func TestPlaceHeadEdge(t *testing.T) {
	q42 := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{42}}
	p813 := Record{Kind: KindReference, Codec: CodecWikidataPID, Payload: []Word{813}}
	st := &docState{}
	if placeHeadEdge(st, headEdgeFixture(EdgeQualifier, 2, p813, EncodeLocalRef(1), q42)) {
		t.Fatal("QUALIFIER before its SOURCE")
	}
	if !placeHeadEdge(st, headEdgeFixture(EdgeSource, 1, EncodeLocalRef(0), q42)) || !st.sources[1] {
		t.Fatalf("SOURCE on the document: %+v", st)
	}
	if !placeHeadEdge(st, headEdgeFixture(EdgeQualifier, 2, p813, EncodeLocalRef(1), q42)) || st.sources[2] {
		t.Fatalf("QUALIFIER on it: %+v", st)
	}
	if placeHeadEdge(st, headEdgeFixture(EdgeQualifier, 3, p813, EncodeLocalRef(2), q42)) {
		t.Fatal("a QUALIFIER on a QUALIFIER is not a HEAD Edge")
	}
	if len(st.doc.Head) != 2 {
		t.Fatalf("head: %d", len(st.doc.Head))
	}
}
