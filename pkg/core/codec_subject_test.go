//ff:func feature=core type=decoder control=sequence
//ff:what TestCodecSubject — a lone TAG WORD is its codec; records with a payload or an Issuer are not
package core

import "testing"

func TestCodecSubject(t *testing.T) {
	r, _, _ := DecodeRecord([]Word{0x0980})
	if w, ok := CodecSubject(r); !ok || w != 0x0980 {
		t.Fatalf("got %04X %v", w, ok)
	}
	if _, ok := CodecSubject(EncodeLocalRef(1)); ok {
		t.Fatal("payload: want false")
	}
	if _, ok := CodecSubject(Record{Kind: KindValue, Codec: CodecValueIssuer2Word, Issuer: []Word{0, 1}}); ok {
		t.Fatal("issuer: want false")
	}
}
