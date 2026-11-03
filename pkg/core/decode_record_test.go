//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeRecord — round trips across Length forms plus prefix, truncation and ISSUER extension errors
package core

import "testing"

func TestDecodeRecord(t *testing.T) {
	for _, n := range []int{0, 13, 14, 0xFFFD, 0xFFFE} {
		in := Record{Kind: KindValue, Codec: CodecUTF8TextEven, Payload: make([]Word, n)}
		ws, err := EncodeRecord(in)
		if err != nil {
			t.Fatal(err)
		}
		got, used, err := DecodeRecord(ws)
		if err != nil || used != len(ws) || len(got.Payload) != n || got.Codec != CodecUTF8TextEven || got.Kind != KindValue {
			t.Fatalf("n=%d: used=%d err=%v", n, used, err)
		}
	}
	got, _, err := DecodeRecord([]Word{0x1BD5, 0, 7, 0, 0xC005, 9})
	if err != nil || got.Kind != KindEdge || got.Codec != CodecEdgeIssuer2Word || got.Issuer[1] != 7 || got.Format != 0xC005 || got.Payload[0] != 9 {
		t.Fatalf("issuer decode: %+v %v", got, err)
	}
	bad := [][]Word{
		{},
		{0x2000},
		{0x1BD3, 0, 7, 0},
		{0x0812},
		{0x03D3, 0, 1, 0},
	}
	for _, b := range bad {
		if _, _, err := DecodeRecord(b); err == nil {
			t.Fatalf("expected error for %04X", b)
		}
	}
	if _, _, err := DecodeRecord([]Word{0x080E, 3, 0}); err == nil {
		t.Fatal("non-canonical 1-word Length extension must fail")
	}
	if _, _, err := DecodeRecord([]Word{0x080F, 0, 2, 0}); err == nil {
		t.Fatal("non-canonical 2-word Length extension must fail")
	}
}
