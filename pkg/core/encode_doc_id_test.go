//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDocID — the DOC-POINTER payload words with version 0; bad widths fail
package core

import (
	"slices"
	"testing"
)

func TestEncodeDocID(t *testing.T) {
	r, err := EncodeDocID([]Word{0, 0}, []Word{0xffff, 0, 0, 1})
	if ws, _ := EncodeRecord(r); err != nil || !slices.Equal(ws, []Word{0x1C79, 0x0418, 0, 0, 0xffff, 0, 0, 1, 0, 0}) {
		t.Fatalf("got %04X %v", ws, err)
	}
	if _, err := EncodeDocID([]Word{0, 0, 0}, []Word{0, 0, 0, 1}); err == nil {
		t.Fatal("issuer: want error")
	}
	if _, err := EncodeDocID([]Word{0, 0}, []Word{1}); err == nil {
		t.Fatal("doc: want error")
	}
}
