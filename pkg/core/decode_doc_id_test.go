//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeDocID — round trip; other records, junk, trailing words, LOCAL, the old EXTERNAL-ISSUER form, a version and a non-Identified Quantifier fail; document 0 decodes
package core

import (
	"slices"
	"testing"
)

func TestDecodeDocID(t *testing.T) {
	r, _ := EncodeDocID([]Word{1, 0x20}, []Word{0xffff, 0, 0, 1})
	if i, d, err := DecodeDocID(r); err != nil || !slices.Equal(i, []Word{1, 0x20}) || !slices.Equal(d, []Word{0xffff, 0, 0, 1}) {
		t.Fatalf("got %v %v %v", i, d, err)
	}
	meta := func(p ...Word) Record { return Record{Kind: KindMeta, Codec: MetaDocID, Payload: p} }
	if _, d, err := DecodeDocID(meta(0x0048, 0, 0, 0, 0, 0, 0, 0, 0)); err != nil || d[3] != 0 {
		t.Fatalf("document 0: %v", err)
	}
	for i, b := range []Record{
		EncodeUintValue(1),
		meta(0, 0, 0, 0, 1),
		meta(append(slices.Clone(r.Payload), 9)...),
		meta(0x0051, 1),
		meta(0x02D7, 0, 0, 0xffff, 0, 0, 1, 0),
		meta(0x0048, 0, 0, 0xffff, 0, 0, 1, 0, 1),
		meta(0x0047, 0, 0xffff, 0, 0, 1, 0, 0),
		meta(0x03C6, 0, 0, 0, 0x0805, 0, 0),
	} {
		if _, _, err := DecodeDocID(b); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
