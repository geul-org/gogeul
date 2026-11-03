//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestTargetPointer — a 4-word ID makes a DOC-POINTER, a 2-word ID a CODEC-POINTER; both decode back; a version list record round trips with either
package core

import (
	"reflect"
	"slices"
	"testing"
)

func TestTargetPointer(t *testing.T) {
	for _, id := range [][]Word{{0, 0, 0, 5}, {0, 0x0980}} {
		r, err := EncodeTargetPointer([]Word{0, 1}, id, []Word{0, 2})
		if err != nil {
			t.Fatal(err)
		}
		iss, gid, ver, err := DecodeTargetPointer(r)
		if err != nil || !slices.Equal(iss, []Word{0, 1}) || !slices.Equal(gid, id) || !slices.Equal(ver, []Word{0, 2}) {
			t.Fatalf("%04X: %04X %04X %04X %v", id, iss, gid, ver, err)
		}
		p := Publication{Codec: EdgeDeprecation, EdgeID: 2, Issuer: []Word{0, 1}, Doc: id, Version: 3}
		if got, err := DecodePublication(EncodePublication(p)); err != nil || !reflect.DeepEqual(got, p) {
			t.Fatalf("publication %04X: %+v %v", id, got, err)
		}
	}
	if _, err := EncodeTargetPointer([]Word{0, 1}, []Word{1, 2, 3}, []Word{0, 0}); err == nil {
		t.Fatal("3-word ID accepted")
	}
	if _, _, _, err := DecodeTargetPointer(Record{Kind: KindReference, Codec: CodecCodecPointer, Payload: []Word{0, 1}}); err == nil {
		t.Fatal("short CODEC-POINTER accepted")
	}
}
