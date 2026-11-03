//ff:func feature=core type=helper control=selection
//ff:what TestFieldAccepts — named slot markers fit only an exact target; wildcards and an empty list refuse them; a Codec 0000 Reference with payload never fits; other records follow the list
package core

import "testing"

func TestFieldAccepts(t *testing.T) {
	marker := func(w Word) Record { r, _, _ := DecodeRecord([]Word{w}); return r }
	object := []Word{0x000F, 0x060F, UnknownValue, NoValue}
	q5 := Record{Kind: KindReference, Codec: CodecWikidataQID, Payload: []Word{5}}
	switch {
	case !FieldAccepts(object, marker(UnknownValue)), !FieldAccepts(object, marker(NoValue)):
		t.Fatal("named markers: want true")
	case FieldAccepts(object, marker(Absent)), FieldAccepts([]Word{0x000F}, marker(UnknownValue)):
		t.Fatal("marker not named, wildcard only: want false")
	case FieldAccepts(nil, marker(Absent)), FieldAccepts(nil, marker(NoValue)):
		t.Fatal("no CODEC list takes no marker")
	case FieldAccepts([]Word{UnknownValue}, Record{Kind: KindReference, Codec: 1, Payload: []Word{7}}):
		t.Fatal("marker C with payload: want false")
	case !FieldAccepts(nil, q5), !FieldAccepts(object, q5), FieldAccepts([]Word{0x0050}, q5):
		t.Fatal("ordinary records follow the list")
	case !FieldAccepts([]Word{0x0680, Absent}, marker(Absent)):
		t.Fatal("optional field: want true")
	}
}
