//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeRecord — END encodes as 0x1C40; an EDGE-ISSUER record carries a 2-word Issuer and a 2-word codec number; codec beyond its zone, Issuer width and codec number on a direct codec fail
package core

import (
	"slices"
	"testing"
)

func TestEncodeRecord(t *testing.T) {
	got, err := EncodeRecord(Record{Kind: KindMeta, Codec: MetaEnd})
	if err != nil || !slices.Equal(got, []Word{0x1C40}) {
		t.Fatalf("END: got %04X err=%v", got, err)
	}
	got, err = EncodeRecord(Record{Kind: KindEdge, Codec: CodecEdgeIssuer2Word, Issuer: []Word{0, 7}, Format: 0xC001, Payload: []Word{1}})
	if err != nil || !slices.Equal(got, []Word{0x1BD5, 0, 7, 0, 0xC001, 1}) {
		t.Fatalf("EDGE-ISSUER: got %04X err=%v", got, err)
	}
	if _, err := EncodeRecord(Record{Kind: KindEdge, Codec: 0x140}); err == nil {
		t.Fatal("Edge codec beyond 0x13F must fail")
	}
	if _, err := EncodeRecord(Record{Kind: KindValue, Codec: 0x40}); err == nil {
		t.Fatal("Value codec beyond 0x3F must fail")
	}
	if _, err := EncodeRecord(Record{Kind: Kind(4)}); err == nil {
		t.Fatal("unknown Kind must fail")
	}
	if _, err := EncodeRecord(Record{Kind: KindEdge, Codec: CodecEdgeIssuer4Word, Issuer: []Word{0, 0}}); err == nil {
		t.Fatal("issuer width mismatch must fail")
	}
	if _, err := EncodeRecord(Record{Kind: KindEdge, Codec: 0x60, Format: 1}); err == nil {
		t.Fatal("codec number on a direct codec must fail")
	}
}
