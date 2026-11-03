//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestTagBoundary — migration.md §7.3 plan2 boundaries by TAG ID: Kind, Codec, ISSUER width and core name at the Kind, Reference section, Value and Edge edges
package core

import "testing"

func TestTagBoundary(t *testing.T) {
	cases := []struct {
		code  Word
		k     Kind
		codec uint16
		iw    int
		name  string
	}{
		{0x05F, KindReference, 0x5F, 16, "REFERENCE-ISSUER-16WORD"}, {0x060, KindValue, 0x00, 0, "BOOL"},
		{0x09F, KindValue, 0x3F, 16, "VALUE-ISSUER-16WORD"}, {0x0A0, KindEdge, 0x00, 0, "REFERENCE-NODE"},
		{0x1BF, KindEdge, 0x11F, 16, "EDGE-ISSUER-16WORD"}, {0x1C0, KindMeta, 0x00, 0, "DOC-1WORD"},
		{0x01D, KindReference, 0x1D, 0, "WIKIDATA-PID"}, {0x01E, KindReference, 0x1E, 0, ""}, {0x01F, KindReference, 0x1F, 0, ""},
		{0x020, KindReference, 0x20, 0, "EXTERNAL-POINTER-1WORD"}, {0x03D, KindReference, 0x3D, 0, "EXTERNAL-ISSUER-16WORD"},
		{0x03E, KindReference, 0x3E, 0, ""}, {0x040, KindReference, 0x40, 0, "ISSUER-POINTER"}, {0x042, KindReference, 0x42, 0, "CODEC-POINTER"},
		{0x043, KindReference, 0x43, 0, ""}, {0x05C, KindReference, 0x5C, 0, ""}, {0x05D, KindReference, 0x5D, 2, "REFERENCE-ISSUER-2WORD"},
		{0x07D, KindValue, 0x1D, 0, "SHA1-DIGEST"}, {0x081, KindValue, 0x21, 0, "DURATION-NANOSECONDS"}, {0x082, KindValue, 0x22, 0, ""},
		{0x09C, KindValue, 0x3C, 0, ""}, {0x09D, KindValue, 0x3D, 2, "VALUE-ISSUER-2WORD"},
		{0x11B, KindEdge, 0x7B, 0, ""}, {0x11C, KindEdge, 0x7C, 0, ""}, {0x11F, KindEdge, 0x7F, 0, ""}, {0x120, KindEdge, 0x80, 0, ""},
		{0x134, KindEdge, 0x94, 0, ""}, {0x135, KindEdge, 0x95, 0, ""}, {0x1BC, KindEdge, 0x11C, 0, ""}, {0x1BD, KindEdge, 0x11D, 2, "EDGE-ISSUER-2WORD"},
	}
	for i := range 3 * len(cases) {
		c, l := cases[i/3], []Word{0, 0xE, 0xF}[i%3]
		k, codec := SplitTarget(c.code<<4 | l)
		if k != c.k || codec != c.codec || IssuerWidth(k, codec) != c.iw || CodecName(k, codec) != c.name {
			t.Fatalf("0x%03XL L=%X: %v 0x%03X %d %q", c.code, l, k, codec, IssuerWidth(k, codec), CodecName(k, codec))
		}
	}
}
