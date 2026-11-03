//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeDeclare — direct targets, ISSUER codes alone and Issuer codecs round-trip; other records, a missing target, a target with L≠0 or prefix ≠ 000, short Issuer and trailing words fail
package core

import (
	"reflect"
	"slices"
	"testing"
)

func TestDecodeDeclare(t *testing.T) {
	for _, d := range []Declare{
		{EdgeID: 1, Target: TargetWord(KindEdge, EdgePublication)},
		{EdgeID: 5, Target: TargetWord(KindValue, CodecValueIssuer2Word), Issuer: []Word{0, 7}, Format: 0xC001},
		{EdgeID: 6, Target: TargetWord(KindReference, CodecRefIssuer4Word), Issuer: []Word{0, 1, 0, 0}, Format: 0x00010002},
		{EdgeID: 7, Target: TargetWord(KindEdge, CodecEdgeIssuer16Word), Issuer: append(make([]Word, 11), 1, 0, 0, 0, 0), Format: 0xC000},
		{EdgeID: 8, Target: TargetWord(KindReference, CodecRefIssuer2Word)},
		{EdgeID: 9, Target: TargetWord(KindEdge, CodecEdgeIssuer16Word)},
	} {
		got, err := DecodeDeclare(EncodeDeclare(d))
		if err != nil || !reflect.DeepEqual(got, d) {
			t.Fatalf("%+v: %+v %v", d, got, err)
		}
	}
	id := localWords(1)
	bad := []Record{
		EncodeUintValue(1),
		{Kind: KindEdge, Codec: EdgeDeclare},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: id},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: slices.Concat(id, []Word{TargetWord(KindEdge, CodecEdgeIssuer2Word), 0, 1, 0})},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: slices.Concat(id, []Word{0x0B80, 0x0051, 2})},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: slices.Concat(id, []Word{0x0AA5})},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: slices.Concat(id, []Word{0x2AA0})},
		{Kind: KindEdge, Codec: EdgeDeclare, Payload: slices.Concat(id, []Word{TargetWord(KindValue, CodecValueIssuer2Word), 0, 7, 0, 0xC001, 7})},
	}
	for i, r := range bad {
		if _, err := DecodeDeclare(r); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
