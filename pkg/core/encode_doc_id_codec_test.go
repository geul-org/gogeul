//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestEncodeDocIDCodec — a codec document's DOC-ID is CODEC-POINTER(Issuer, codec ID, 0); DOC-POINTER, a version and other records are rejected
package core

import (
	"slices"
	"testing"
)

func TestEncodeDocIDCodec(t *testing.T) {
	r, err := EncodeDocIDCodec([]Word{0, 0}, 0x00000980)
	ws, _ := EncodeRecord(r)
	if err != nil || !slices.Equal(ws, []Word{0x1C77, 0x03C6, 0, 0, 0, 0x0980, 0, 0}) {
		t.Fatalf("%04X %v", ws, err)
	}
	if iss, c, err := DecodeDocIDCodec(r); err != nil || !slices.Equal(iss, []Word{0, 0}) || c != 0x980 {
		t.Fatalf("%v %x %v", iss, c, err)
	}
	doc, _ := EncodeDocID([]Word{0, 0}, []Word{0, 0, 0, 1})
	ver, _ := EncodeRecord(Record{Kind: KindReference, Codec: CodecCodecPointer, Payload: []Word{0, 0, 0, 0x0980, 0, 1}})
	for i, b := range []Record{doc, EncodeUintValue(1), {Kind: KindMeta, Codec: MetaDocID, Payload: ver}, {Kind: KindMeta, Codec: MetaDocID, Payload: append(slices.Clone(r.Payload), 1)}} {
		if _, _, err := DecodeDocIDCodec(b); err == nil {
			t.Fatalf("bad %d accepted", i)
		}
	}
	if _, err := EncodeDocIDCodec([]Word{0}, 1); err == nil {
		t.Fatal("1-word Issuer accepted")
	}
}
