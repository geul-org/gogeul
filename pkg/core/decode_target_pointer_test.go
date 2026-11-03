//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestDecodeTargetPointer — a CODEC-POINTER gives a 2-word ID, a DOC-POINTER a 4-word ID; a broken CODEC-POINTER and other codes are errors
package core

import (
	"slices"
	"testing"
)

func TestDecodeTargetPointer(t *testing.T) {
	cases := []struct {
		ws      []Word
		id, ver []Word
	}{
		{[]Word{0x03C6, 0, 1, 0, 0x0980, 0, 2}, []Word{0, 0x0980}, []Word{0, 2}},
		{[]Word{0x0048, 0, 1, 0, 0, 0, 5, 0, 2}, []Word{0, 0, 0, 5}, []Word{0, 2}},
	}
	for _, c := range cases {
		r, _, err := DecodeRecord(c.ws)
		if err != nil {
			t.Fatal(err)
		}
		iss, id, ver, err := DecodeTargetPointer(r)
		if err != nil || !slices.Equal(iss, []Word{0, 1}) || !slices.Equal(id, c.id) || !slices.Equal(ver, c.ver) {
			t.Fatalf("%04X: %04X %04X %04X %v", c.ws, iss, id, ver, err)
		}
	}
	for _, bad := range []Record{
		{Kind: KindReference, Codec: CodecCodecPointer, Payload: []Word{0, 1}},
		{Kind: KindReference, Codec: CodecIssuerPointer, Payload: []Word{0, 1}},
	} {
		if _, _, _, err := DecodeTargetPointer(bad); err == nil {
			t.Fatalf("%+v accepted", bad)
		}
	}
}
