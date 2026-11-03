//ff:func feature=core type=decoder control=iteration dimension=2
//ff:what TestEncodeCodecPointer — CODEC-POINTER 0x03CL round trips with 2/4/16-word Issuers × version 0·1; payload 6/8/20 only, not ISSUER nor quantified
package core

import (
	"slices"
	"testing"
)

func TestEncodeCodecPointer(t *testing.T) {
	issuers := [][]Word{{0, 0}, {0, 1, 0, 0}, append(make([]Word, 11), 1, 0, 0, 0, 0)}
	for _, iss := range issuers {
		for _, v := range []uint32{0, 1} {
			r, err := EncodeCodecPointer(iss, 0x000003D0, v)
			ws, _ := EncodeRecord(r)
			if err != nil || ws[0] != 0x03C0|Word(len(iss)+4) && len(iss) < 16 {
				t.Fatalf("%d-word: %04X %v", len(iss), ws, err)
			}
			got, n, _ := DecodeRecord(ws)
			gi, gc, gv, err := DecodeCodecPointer(got)
			if err != nil || n != len(ws) || !slices.Equal(gi, iss) || gc != 0x3D0 || gv != v {
				t.Fatalf("%d-word v%d: %04X %x %d %v", len(iss), v, gi, gc, gv, err)
			}
		}
	}
	if IssuerWidth(KindReference, CodecCodecPointer) != 0 || QuantifierName(CodecCodecPointer) != "" || IsQuantified(CodecCodecPointer) {
		t.Fatal("CODEC-POINTER read as ISSUER or quantified")
	}
	for _, p := range [][]Word{{0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0}, {0, 0, 0, 1, 0, 0, 0, 0}, make([]Word, 10), make([]Word, 20)} {
		if _, _, _, err := DecodeCodecPointer(Record{Kind: KindReference, Codec: CodecCodecPointer, Payload: p}); err == nil {
			t.Fatalf("%d-word payload %04X accepted", len(p), p)
		}
	}
	if _, _, _, err := DecodeCodecPointer(Record{Kind: KindReference, Codec: CodecDocPointer, Payload: make([]Word, 6)}); err == nil {
		t.Fatal("DOC-POINTER accepted")
	}
	if _, err := EncodeCodecPointer([]Word{1}, 1, 0); err == nil {
		t.Fatal("1-word Issuer accepted")
	}
}
