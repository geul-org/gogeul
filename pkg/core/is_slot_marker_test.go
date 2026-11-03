//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsSlotMarker — 0x0000, 0x0010, 0x0020 are one-word markers; pointers, LOCAL and payloads are not
package core

import "testing"

func TestIsSlotMarker(t *testing.T) {
	for _, w := range []Word{Absent, UnknownValue, NoValue} {
		r, n, err := DecodeRecord([]Word{w, 0x1234})
		if err != nil || n != 1 || !IsSlotMarker(r) {
			t.Fatalf("0x%04X: %+v %d %v", w, r, n, err)
		}
	}
	for _, r := range []Record{EncodeLocalRef(1), {Kind: KindReference, Codec: 1, Payload: []Word{1}}, {Kind: KindValue, Codec: 1},
		{Kind: KindReference, Codec: CodecLocal}, {Kind: KindReference, Codec: CodecIssuerPointer}, {Kind: KindReference, Codec: CodecDocPointer}} {
		if IsSlotMarker(r) {
			t.Fatalf("%+v: want false", r)
		}
	}
}
