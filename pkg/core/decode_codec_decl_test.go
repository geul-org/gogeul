//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeCodecDecl — a round trip; another Edge, broken LOCALs, a missing or extra word, LLLL bits and an empty wildcard fail
package core

import (
	"slices"
	"testing"
)

func TestDecodeCodecDecl(t *testing.T) {
	in := CodecDecl{EdgeID: 3, Field: 2, Target: TargetWord(KindReference, CodecWikidataQID)}
	if got, err := DecodeCodecDecl(EncodeCodecDecl(in)); err != nil || got != in {
		t.Fatalf("%+v %v", got, err)
	}
	id := localWords(3)
	for i, r := range []Record{
		{Kind: KindEdge, Codec: EdgeEnum},
		{Kind: KindEdge, Codec: EdgeCodec},
		{Kind: KindEdge, Codec: EdgeCodec, Payload: slices.Concat(id, []Word{1})},
		{Kind: KindEdge, Codec: EdgeCodec, Payload: slices.Concat(id, localWords(2))},
		{Kind: KindEdge, Codec: EdgeCodec, Payload: slices.Concat(id, localWords(2), []Word{0x084A, 0})},
		{Kind: KindEdge, Codec: EdgeCodec, Payload: slices.Concat(id, localWords(2), []Word{0x0D04})},
		{Kind: KindEdge, Codec: EdgeCodec, Payload: slices.Concat(id, localWords(2), []Word{0x0F84})},
	} {
		if _, err := DecodeCodecDecl(r); err == nil {
			t.Fatalf("case %d: want error", i)
		}
	}
}
