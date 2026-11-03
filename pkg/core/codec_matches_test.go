//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestCodecMatches — exact targets (C 0 included), Kind wildcards over every KKK, Reference enumeration; slot markers only by exact target
package core

import "testing"

func TestCodecMatches(t *testing.T) {
	cases := []struct {
		t     Word
		k     Kind
		codec uint16
		want  bool
	}{
		{0x0140, KindReference, CodecWikidataQID, true},
		{0x0140, KindReference, CodecWikidataQID + 3, false},
		{0x0170, KindReference, CodecWikidataQID + 3, true},
		{0x000F, KindReference, CodecWikidataQID + 4, true},
		{0x000F, KindReference, CodecRefIssuer2Word, true},
		{0x000F, KindReference, CodecCodecPointer, true},
		{0x014F, KindReference, CodecLocal, true},
		{0x000F, KindValue, CodecUnsignedInteger, false},
		{0x000F, KindReference, 0, false},
		{0x000F, KindReference, 1, false},
		{0x000F, KindReference, 2, false},
		{0x000F, KindReference, CodecIssuerPointer, true},
		{0x040F, KindValue, CodecCoordinateEarth, true},
		{0x047F, KindValue, CodecValueIssuer16Word, true},
		{0x040F, KindEdge, EdgeBind, false},
		{0x0400, KindValue, CodecBool, true},
		{0x0400, KindValue, CodecSignedInteger, false},
		{0x0580, KindValue, CodecUTF8LanguageTextEven, true},
		{0x0580, KindValue, CodecUTF8LanguageTextEven | 1, false},
		{0x0800, KindEdge, 0, true},
		{0x0800, KindEdge, EdgeBind, false},
		{0x080F, KindEdge, CodecEdgeIssuer2Word, true},
		{0x0C0F, KindEdge, EdgeTriple, true},
		{0x100F, KindEdge, EdgeTriple, true},
		{0x140F, KindEdge, EdgeTriple, true},
		{0x180F, KindEdge, EdgeTriple, true},
		{0x1BDF, KindEdge, 0, true},
		{0x1BD0, KindEdge, CodecEdgeIssuer2Word, true},
		{0x1BD0, KindEdge, CodecEdgeIssuer4Word, false},
		{0x180F, KindMeta, MetaBody, false},
		{0x1C00, KindMeta, MetaDoc1Word, true},
		{0x1C00, KindMeta, MetaBody, false},
		{0x1C0F, KindMeta, MetaVerdict, true},
		{0x1FFF, KindMeta, MetaBody, true},
		{0x1C0F, KindEdge, EdgeTriple, false},
		{Absent, KindReference, 0, true},
		{UnknownValue, KindReference, 1, true},
		{NoValue, KindReference, 2, true},
		{UnknownValue, KindReference, 2, false},
		{NoValue, KindReference, 0, false},
		{UnknownValue, KindValue, 1, false},
		{0x0030, KindReference, CodecIssuerPointer, true},
		{0x0040, KindReference, CodecDocPointer, true},
		{0x03C0, KindReference, CodecCodecPointer, true},
	}
	for i, c := range cases {
		if got := CodecMatches(c.t, c.k, c.codec); got != c.want {
			t.Fatalf("case %d (%04X, %v, 0x%03X): %v", i, c.t, c.k, c.codec, got)
		}
	}
}
