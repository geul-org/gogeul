//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeCodecPointer — CODEC-POINTER 0x03CL payload splits into Issuer, codec ID and version ID; other codes, short payloads and non-canonical Issuers are rejected
package core

import (
	"slices"
	"testing"
)

func TestDecodeCodecPointer(t *testing.T) {
	ws := []Word{0x03C6, 0, 1, 0, 0x0980, 0, 2}
	r, n, err := DecodeRecord(ws)
	if err != nil || n != 7 || r.Codec != CodecCodecPointer {
		t.Fatalf("%+v %d %v", r, n, err)
	}
	iss, c, v, err := DecodeCodecPointer(r)
	if err != nil || !slices.Equal(iss, []Word{0, 1}) || c != 0x0980 || v != 2 {
		t.Fatalf("%04X %x %d %v", iss, c, v, err)
	}
	for _, bad := range []Record{
		{Kind: KindValue, Codec: CodecCodecPointer, Payload: make([]Word, 6)},
		{Kind: KindReference, Codec: CodecDocPointer, Payload: make([]Word, 6)},
		{Kind: KindReference, Codec: CodecCodecPointer, Payload: make([]Word, 5)},
		{Kind: KindReference, Codec: CodecCodecPointer, Payload: []Word{0, 0, 0, 1, 0, 0, 0, 0}},
	} {
		if _, _, _, err := DecodeCodecPointer(bad); err == nil {
			t.Fatalf("%+v accepted", bad)
		}
	}
}
