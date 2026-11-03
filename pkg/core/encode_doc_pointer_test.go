//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what TestEncodeDocPointer — a 2-word Issuer document version encodes as 0048 …; document 0 (the term definition index) encodes; bad widths fail
package core

import (
	"slices"
	"testing"
)

func TestEncodeDocPointer(t *testing.T) {
	doc := []Word{0xFFFF, 0, 0, 1}
	r, err := EncodeDocPointer([]Word{0, 0}, doc, []Word{0, 1})
	if err != nil {
		t.Fatal(err)
	}
	if ws, _ := EncodeRecord(r); !slices.Equal(ws, []Word{0x0048, 0, 0, 0xFFFF, 0, 0, 1, 0, 1}) {
		t.Fatalf("%04X", ws)
	}
	if _, err := EncodeDocPointer([]Word{0, 0}, []Word{0, 0, 0, 0}, []Word{0, 0}); err != nil {
		t.Fatalf("document 0: %v", err)
	}
	for _, c := range [][3][]Word{
		{{1, 2, 3}, doc, {0, 1}},
		{{0}, doc, {0, 1}},
		{{0, 0}, {1, 2, 3}, {0, 1}},
		{{0, 0}, doc, {1}},
	} {
		if _, err := EncodeDocPointer(c[0], c[1], c[2]); err == nil {
			t.Fatalf("%v accepted", c)
		}
	}
}
