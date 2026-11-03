//ff:func feature=core type=helper control=sequence
//ff:what TestIsAbsent — WORD 0x0000 decodes as the one-word Absent record; LOCAL and a Codec-0 Reference with payload are not
package core

import "testing"

func TestIsAbsent(t *testing.T) {
	r, n, err := DecodeRecord([]Word{Absent, 0x1234})
	if err != nil || n != 1 || !IsAbsent(r) {
		t.Fatalf("0x0000: %+v %d %v", r, n, err)
	}
	if IsAbsent(EncodeLocalRef(1)) || IsAbsent(Record{Kind: KindReference, Payload: []Word{1}}) || IsAbsent(Record{Kind: KindValue}) {
		t.Fatal("non-absent records: want false")
	}
	if !CodecMatches(Absent, r.Kind, r.Codec) || CodecMatches(Absent, KindReference, CodecLocal) {
		t.Fatal("CODEC 0x0000 matches only Absent")
	}
}
