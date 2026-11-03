//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodePublication — the three codecs decode; wrong kind, bad Edge Id, a missing or bad DOC-POINTER, wrong tail length and version 0 fail
package core

import "testing"

func TestDecodePublication(t *testing.T) {
	id, _ := EncodeRecord(EncodeLocalRef(3))
	ptr := func(v Word) []Word { return []Word{0x0418, 0, 0, 0, 0, 0, 5, 0, v} }
	edge := func(c uint16, rest ...Word) Record {
		return Record{Kind: KindEdge, Codec: c, Payload: append(append([]Word{}, id...), rest...)}
	}
	pub := edge(EdgePublication, append(ptr(2), make([]Word, 16)...)...)
	if p, err := DecodePublication(pub); err != nil || p.EdgeID != 3 || len(p.Issuer) != 2 || p.Issuer[1] != 0 || p.Doc[3] != 5 || p.Version != 2 {
		t.Fatalf("publication: %+v %v", p, err)
	}
	if p, err := DecodePublication(edge(EdgeTombstone, append(ptr(1), 9)...)); err != nil || p.Reason != 9 {
		t.Fatalf("tombstone: %+v %v", p, err)
	}
	bad := []Record{
		{Kind: KindValue, Codec: EdgePublication},
		{Kind: KindEdge, Codec: EdgeDeprecation},
		{Kind: KindEdge, Codec: EdgeDeprecation, Payload: append([]Word{0x0621, 1}, ptr(1)...)},
		edge(EdgeDeprecation),
		edge(EdgeDeprecation, 0, 0, 0, 0, 5, 0, 1),
		edge(EdgeDeprecation, append(ptr(1), 7)...),
		edge(EdgeDeprecation, ptr(0)...),
	}
	for i, r := range bad {
		if _, err := DecodePublication(r); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
