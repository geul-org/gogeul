//ff:func feature=core type=encoder control=sequence
//ff:what TestTagWord — the TAG is the Kind zone start + Codec<<4 + L: DOC-1WORD 0x1C01, REFERENCE-ISSUER-16WORD 0x03F0, EDGE-ISSUER-16WORD L=15 0x1BFF
package core

import "testing"

func TestTagWord(t *testing.T) {
	if got := tagWord(KindMeta, 1, MetaDoc1Word); got != 0x1C01 {
		t.Fatalf("DOC-1WORD tag: got 0x%04X", uint16(got))
	}
	if got := tagWord(KindReference, 0, 0x3F); got != 0x03F0 {
		t.Fatalf("reference tag: got 0x%04X", uint16(got))
	}
	if got := tagWord(KindEdge, 0xF, 0x13F); got != 0x1BFF {
		t.Fatalf("edge issuer tag: got 0x%04X", uint16(got))
	}
	if got := tagWord(KindValue, 2, CodecUTF8LanguageTextEven); got != 0x0582 {
		t.Fatalf("value tag: got 0x%04X", uint16(got))
	}
}
