//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeFieldDecl — raw and record fields round-trip with their DECLARE; other records, a missing DECLARE and bad layouts fail
package core

import (
	"slices"
	"testing"
)

func TestDecodeFieldDecl(t *testing.T) {
	for _, f := range []FieldDecl{{EdgeID: 2, Declare: 1, Width: 4}, {EdgeID: 3, Declare: 1}} {
		got, err := DecodeFieldDecl(EncodeFieldDecl(f))
		if err != nil || got != f {
			t.Fatalf("%+v: %+v %v", f, got, err)
		}
	}
	id := localWords(1)
	for i, r := range []Record{
		EncodeUintValue(1),
		{Kind: KindEdge, Codec: EdgeField},
		{Kind: KindEdge, Codec: EdgeField, Payload: id},
		{Kind: KindEdge, Codec: EdgeField, Payload: slices.Concat(id, []Word{4})},
		{Kind: KindEdge, Codec: EdgeField, Payload: slices.Concat(id, id)},
		{Kind: KindEdge, Codec: EdgeField, Payload: slices.Concat(id, id, []Word{4, 0x0051})},
		{Kind: KindEdge, Codec: EdgeField, Payload: slices.Concat(id, id, []Word{0, 0x0051})},
	} {
		if _, err := DecodeFieldDecl(r); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
