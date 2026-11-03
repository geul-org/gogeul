//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestCodeName — slot markers, 5Q codes of every scheme and Quantifier, even/odd pair members, unpaired and unassigned codes, TOULMIN codes, other Kinds as CodecName
package core

import "testing"

func TestCodeName(t *testing.T) {
	cases := []struct {
		k    Kind
		c    uint16
		want string
	}{
		{KindReference, 0, "ABSENT"}, {KindReference, 1, "UNKNOWN-VALUE"}, {KindReference, 2, "NO-VALUE"},
		{KindReference, 3, ""}, {KindReference, 4, ""},
		{KindReference, 0x05, "LOCAL-IDENTIFIED"}, {KindReference, 0x06, "LOCAL-GENERIC"}, {KindReference, 0x07, "LOCAL-FREE-CHOICE"},
		{KindReference, 0x08, "LOCAL-UNIVERSAL"}, {KindReference, 0x09, "LOCAL-VARIABLE"},
		{KindReference, 0x0A, "GEUL-IDENTIFIED"}, {KindReference, 0x10, "GEUL-SIDX-GENERIC"}, {KindReference, 0x18, "WIKIDATA-QID-VARIABLE"},
		{KindReference, 0x19, "WIKIDATA-PID-IDENTIFIED"}, {KindReference, 0x1D, "WIKIDATA-PID-VARIABLE"}, {KindReference, 0x1E, ""},
		{KindReference, 0x20, "EXTERNAL-POINTER-1WORD-IDENTIFIED"}, {KindReference, 0x21, "EXTERNAL-POINTER-1WORD-GENERIC"},
		{KindReference, 0x2E, "EXTERNAL-POINTER-4WORD-VARIABLE"}, {KindReference, 0x2F, "EXTERNAL-ISSUER-2WORD-IDENTIFIED"},
		{KindReference, 0x3B, "EXTERNAL-ISSUER-16WORD-FREE-CHOICE"}, {KindReference, 0x3D, "EXTERNAL-ISSUER-16WORD-VARIABLE"}, {KindReference, 0x3E, ""},
		{KindReference, 0x40, "ISSUER-POINTER"}, {KindReference, 0x42, "CODEC-POINTER"}, {KindReference, 0x5F, "REFERENCE-ISSUER-16WORD"},
		{KindValue, 0x0E, "EDTF-EVEN"}, {KindValue, 0x0F, "EDTF-ODD"}, {KindValue, 0x10, "GEDCOM-DATE-EVEN"}, {KindValue, 0x11, "GEDCOM-DATE-ODD"},
		{KindValue, 0x12, "UTF8-TEXT-EVEN"}, {KindValue, 0x13, "UTF8-TEXT-ODD"}, {KindValue, 0x14, "URI-EVEN"}, {KindValue, 0x15, "URI-ODD"},
		{KindValue, 0x16, "LANGUAGE-BCP47-CODE-EVEN"}, {KindValue, 0x17, "LANGUAGE-BCP47-CODE-ODD"},
		{KindValue, 0x18, "UTF8-LANGUAGE-TEXT-EVEN"}, {KindValue, 0x19, "UTF8-LANGUAGE-TEXT-ODD"},
		{KindValue, 0x1E, "OPAQUE-BYTES-EVEN"}, {KindValue, 0x1F, "OPAQUE-BYTES-ODD"},
		{KindValue, 0x00, "BOOL"}, {KindValue, 0x0D, "DURATION-SECONDS"}, {KindValue, 0x1A, "LANGUAGE-SHORT-CODE"}, {KindValue, 0x1D, "SHA1-DIGEST"},
		{KindValue, 0x20, "BIT-VECTOR"}, {KindValue, 0x21, "DURATION-NANOSECONDS"}, {KindValue, 0x22, ""}, {KindValue, 0x3D, "VALUE-ISSUER-2WORD"},
		{KindEdge, EdgeTriple, "TRIPLE"}, {KindEdge, 0x11D, "EDGE-ISSUER-2WORD"}, {KindEdge, 0x60, "CASE-DECLARE"},
		{KindEdge, 0x6A, "SUPERSEDE-CHAIN"}, {KindEdge, 0x80, "SOURCE"}, {KindEdge, 0x7C, ""}, {KindEdge, 0x95, ""},
		{KindMeta, MetaDocID, "DOC-ID"}, {KindMeta, 0x2A, ""},
	}
	for _, c := range cases {
		if got := CodeName(c.k, c.c); got != c.want {
			t.Fatalf("CodeName(%v, 0x%03X)=%q want %q", c.k, c.c, got, c.want)
		}
	}
}
