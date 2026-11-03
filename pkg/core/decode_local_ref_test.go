//ff:func feature=core type=decoder control=sequence
//ff:what TestDecodeLocalRef — round trip, wrong codec and non-minimal ID
package core

import "testing"

func TestDecodeLocalRef(t *testing.T) {
	if id, err := DecodeLocalRef(EncodeLocalRef(1 << 33)); err != nil || id != 1<<33 {
		t.Fatalf("got %d %v", id, err)
	}
	if _, err := DecodeLocalRef(Record{Kind: KindReference, Codec: CodecGeul, Payload: []Word{1}}); err == nil {
		t.Fatal("GEUL Reference must fail")
	}
	if _, err := DecodeLocalRef(Record{Kind: KindReference, Codec: CodecLocal, Payload: []Word{0, 1}}); err == nil {
		t.Fatal("non-minimal ID must fail")
	}
}
