//ff:func feature=core type=decoder control=sequence
//ff:what TestHeadEdge — SOURCE on LOCAL 0 and QUALIFIER on a HEAD SOURCE are HEAD Edges; SOURCE on a record, other QUALIFIERs, other codecs, ISSUER forms, Meta and broken Edges are not
package core

import "testing"

func TestHeadEdge(t *testing.T) {
	q42 := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{42}}
	p813 := Record{Kind: KindReference, Codec: CodecWikidataPID, Payload: []Word{813}}
	if id, ok := HeadEdge(headEdgeFixture(EdgeSource, 3, EncodeLocalRef(0), q42), nil); !ok || id != 3 {
		t.Fatalf("SOURCE on the document: %d %v", id, ok)
	}
	if _, ok := HeadEdge(headEdgeFixture(EdgeSource, 3, EncodeLocalRef(1), q42), nil); ok {
		t.Fatal("SOURCE on a record is a BODY Edge")
	}
	qual := headEdgeFixture(EdgeQualifier, 4, p813, EncodeLocalRef(3), q42)
	if id, ok := HeadEdge(qual, map[uint64]bool{3: true}); !ok || id != 4 {
		t.Fatalf("QUALIFIER on a HEAD SOURCE: %d %v", id, ok)
	}
	if _, ok := HeadEdge(qual, map[uint64]bool{9: true}); ok {
		t.Fatal("QUALIFIER on another Edge is a BODY Edge")
	}
	if _, ok := HeadEdge(headEdgeFixture(EdgeTriple, 5, p813, EncodeLocalRef(0), q42), nil); ok {
		t.Fatal("TRIPLE is never a HEAD Edge")
	}
	issuer := headEdgeFixture(EdgeSource, 3, EncodeLocalRef(0), q42)
	issuer.Issuer = []Word{0, 1}
	if _, ok := HeadEdge(issuer, nil); ok {
		t.Fatal("an ISSUER form is not the direct SOURCE")
	}
	if _, ok := HeadEdge(Record{Kind: KindMeta, Codec: MetaAuthor, Payload: issuer.Payload}, nil); ok {
		t.Fatal("Meta AUTHOR is not an Edge")
	}
	if _, ok := HeadEdge(Record{Kind: KindEdge, Codec: EdgeSource, Payload: []Word{0x0140}}, nil); ok {
		t.Fatal("an Edge without a LOCAL Edge Id")
	}
	if _, ok := HeadEdge(headEdgeFixture(EdgeSource, 3, q42), nil); ok {
		t.Fatal("a SOURCE whose Target is not a LOCAL")
	}
}
