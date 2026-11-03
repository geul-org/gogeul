//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestCodecName — spot names per Kind, CODEC-POINTER, REFERENCE-ISSUER, ISSUER codes per Kind; empty codes, markers and Reserved Edge codes are ""
package core

import "testing"

func TestCodecName(t *testing.T) {
	cases := []struct {
		k    Kind
		c    uint16
		want string
	}{
		{KindMeta, MetaImportDoc, "IMPORT-DOC"}, {KindMeta, MetaEffectiveFrom, "EFFECTIVE-FROM"}, {KindMeta, 0x10, "RECORDED-AT"},
		{KindMeta, 0x25, "ISSUER-KEY"}, {KindMeta, 0x29, "VERDICT"}, {KindMeta, 0x2A, ""}, {KindMeta, 0x3F, ""},
		{KindValue, CodecUTF8TextEven | 1, "UTF8-TEXT"}, {KindValue, 0x0C, "POSIX-TIMESTAMP-SECONDS"}, {KindValue, 0x0F, "EDTF"},
		{KindValue, 0x18, "UTF8-LANGUAGE-TEXT"}, {KindValue, 0x1A, "LANGUAGE-SHORT-CODE"}, {KindValue, 0x1D, "SHA1-DIGEST"},
		{KindValue, 0x1E, "OPAQUE-BYTES"}, {KindValue, 0x1F, "OPAQUE-BYTES"}, {KindValue, 0x20, "BIT-VECTOR"}, {KindValue, 0x21, "DURATION-NANOSECONDS"}, {KindValue, 0x22, ""}, {KindValue, 0x3C, ""}, {KindValue, 0x3D, "VALUE-ISSUER-2WORD"}, {KindValue, 0x3F, "VALUE-ISSUER-16WORD"},
		{KindEdge, EdgeSource, "SOURCE"}, {KindEdge, 0x21, ""}, {KindEdge, 0x41, "FUNCTION-DECLARE"}, {KindEdge, EdgeBind, "BIND"},
		{KindEdge, 0x40, ""}, {KindEdge, 0x7C, ""}, {KindEdge, 0x7F, ""}, {KindEdge, 0x80, ""}, {KindEdge, 0x94, ""}, {KindEdge, 0x95, ""}, {KindEdge, 0x13C, ""},
		{KindEdge, 0x13D, "EDGE-ISSUER-2WORD"}, {KindEdge, 0x13F, "EDGE-ISSUER-16WORD"},
		{KindReference, CodecLocal + 4, "LOCAL"}, {KindReference, 0x37, "EXTERNAL-ISSUER-16WORD"}, {KindReference, 0x3B, "EXTERNAL-ISSUER-16WORD"},
		{KindReference, 0x3C, "CODEC-POINTER"}, {KindReference, 0x3D, "REFERENCE-ISSUER-2WORD"}, {KindReference, 0x3F, "REFERENCE-ISSUER-16WORD"},
		{KindReference, 3, "ISSUER-POINTER"}, {KindReference, 4, "DOC-POINTER"}, {KindReference, 0, ""}, {KindReference, 2, ""},
	}
	for _, c := range cases {
		if got := CodecName(c.k, c.c); got != c.want {
			t.Fatalf("CodecName(%v, 0x%03X)=%q want %q", c.k, c.c, got, c.want)
		}
	}
}
