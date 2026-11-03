//ff:func feature=core type=decoder control=sequence
//ff:what TestLocalAt — the LOCAL at a record index of a payload; false past the end, on a broken record or a non-LOCAL
package core

import "testing"

func TestLocalAt(t *testing.T) {
	src := headEdgeFixture(EdgeSource, 7, EncodeLocalRef(0), Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{42}})
	if id, ok := localAt(src.Payload, 0); !ok || id != 7 {
		t.Fatalf("Edge Id: %d %v", id, ok)
	}
	if id, ok := localAt(src.Payload, 1); !ok || id != 0 {
		t.Fatalf("Target 0: %d %v", id, ok)
	}
	if _, ok := localAt(src.Payload, 2); ok {
		t.Fatal("a WIKIDATA-QID is not a LOCAL")
	}
	if _, ok := localAt(src.Payload, 3); ok {
		t.Fatal("past the end")
	}
	if _, ok := localAt([]Word{0x0051}, 1); ok {
		t.Fatal("a truncated record before the index")
	}
}
