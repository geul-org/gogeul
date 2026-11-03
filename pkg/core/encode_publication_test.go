//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestEncodePublication — each codec round-trips through DecodePublication; the DOC-POINTER words follow the Edge Id
package core

import (
	"reflect"
	"slices"
	"testing"
)

func TestEncodePublication(t *testing.T) {
	doc := []Word{0, 0, 0, 5}
	for _, p := range []Publication{
		{Codec: EdgePublication, EdgeID: 1, Issuer: []Word{0, 0}, Doc: doc, Version: 2, Hash: [32]byte{31: 0xAB}},
		{Codec: EdgeDeprecation, EdgeID: 2, Issuer: []Word{1, 0x20}, Doc: doc, Version: 2},
		{Codec: EdgeTombstone, EdgeID: 3, Issuer: []Word{0, 0}, Doc: doc, Version: 2, Reason: 4},
	} {
		got, err := DecodePublication(EncodePublication(p))
		if err != nil || !reflect.DeepEqual(got, p) {
			t.Fatalf("%+v: %+v %v", p, got, err)
		}
	}
	r := EncodePublication(Publication{Codec: EdgeDeprecation, EdgeID: 1, Issuer: []Word{0, 0}, Doc: doc, Version: 1})
	if !slices.Equal(r.Payload, []Word{0x0051, 1, 0x0048, 0, 0, 0, 0, 0, 5, 0, 1}) {
		t.Fatalf("words %04X", r.Payload)
	}
}
