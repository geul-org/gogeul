//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestCodecMatches — exact targets (C 0 included); section wildcards with C 0 and C ≠ 0, both sides of every section boundary, other sections of the same Kind not matching; slot markers only by exact target
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
		{0x0400, KindReference, CodecIssuerPointer, true},
		{0x0410, KindReference, CodecDocPointer, true},
		{0x0420, KindReference, CodecCodecPointer, true},
		{0x0420, KindReference, CodecDocPointer, false},
		// Reference section 0 wildcard: entities only.
		{0x000F, KindReference, CodecWikidataQID + 4, true},
		{0x014F, KindReference, CodecLocal, true},
		{0x000F, KindReference, 0x1F, true},
		{0x000F, KindReference, CodecExternalPointer1Word, false},
		{0x000F, KindReference, CodecIssuerPointer, false},
		{0x000F, KindReference, 0, false},
		{0x000F, KindReference, 1, false},
		{0x000F, KindReference, 2, false},
		{0x000F, KindValue, CodecUnsignedInteger, false},
		// section 1 (EXTERNAL) and section 2.
		{0x020F, KindReference, CodecExternalPointer1Word, true},
		{0x03DF, KindReference, CodecExternalIssuer16Word + 4, true},
		{0x020F, KindReference, CodecWikidataPID + 4, false},
		{0x020F, KindReference, CodecIssuerPointer, false},
		{0x040F, KindReference, CodecIssuerPointer, true},
		{0x05DF, KindReference, CodecCodecPointer, true},
		{0x040F, KindReference, CodecRefIssuer2Word, true},
		{0x040F, KindReference, CodecExternalIssuer16Word + 4, false},
		// Value sections 3 and 4.
		{0x060F, KindValue, CodecCoordinateEarth, true},
		{0x07FF, KindValue, CodecSHA1Digest, true},
		{0x060F, KindValue, CodecBitVector, false},
		{0x080F, KindValue, CodecBitVector, true},
		{0x09FF, KindValue, CodecValueIssuer16Word, true},
		{0x080F, KindValue, CodecOpaqueBytesEven | 1, false},
		{0x060F, KindEdge, EdgeBind, false},
		{0x0600, KindValue, CodecBool, true},
		{0x0600, KindValue, CodecSignedInteger, false},
		{0x0780, KindValue, CodecUTF8LanguageTextEven, true},
		{0x0780, KindValue, CodecUTF8LanguageTextEven | 1, false},
		// Edge sections 5~13.
		{0x0A00, KindEdge, 0, true},
		{0x0A00, KindEdge, EdgeBind, false},
		{0x0A0F, KindEdge, EdgeTriple, true},
		{0x0A0F, KindEdge, EdgeSource, false},
		{0x0C0F, KindEdge, EdgeSource, true},
		{0x0BFF, KindEdge, EdgeCodec, true},
		{0x0A0F, KindEdge, CodecEdgeIssuer2Word, false},
		{0x1A0F, KindEdge, CodecEdgeIssuer2Word, true},
		{0x1BDF, KindEdge, CodecEdgeIssuer16Word, true},
		{0x1BD0, KindEdge, CodecEdgeIssuer2Word, true},
		{0x1BD0, KindEdge, CodecEdgeIssuer4Word, false},
		{0x1A0F, KindMeta, MetaBody, false},
		// Meta sections 14 and 15.
		{0x1C00, KindMeta, MetaDoc1Word, true},
		{0x1C00, KindMeta, MetaBody, false},
		{0x1C0F, KindMeta, MetaAccess, true},
		{0x1C0F, KindMeta, MetaPrivacy, false},
		{0x1E0F, KindMeta, MetaVerdict, true},
		{0x1FFF, KindMeta, MetaPrivacy, true},
		{0x1FFF, KindMeta, MetaBody, false},
		{0x1C0F, KindEdge, EdgeTriple, false},
		{Absent, KindReference, 0, true},
		{UnknownValue, KindReference, 1, true},
		{NoValue, KindReference, 2, true},
		{UnknownValue, KindReference, 2, false},
		{NoValue, KindReference, 0, false},
		{UnknownValue, KindValue, 1, false},
	}
	for i, c := range cases {
		if got := CodecMatches(c.t, c.k, c.codec); got != c.want {
			t.Fatalf("case %d (%04X, %v, 0x%03X): %v", i, c.t, c.k, c.codec, got)
		}
	}
}
