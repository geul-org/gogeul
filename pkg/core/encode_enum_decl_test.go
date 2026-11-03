//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeEnumDecl — the edge/ENUM example words
package core

import (
	"slices"
	"testing"
)

func TestEncodeEnumDecl(t *testing.T) {
	ws, _ := EncodeRecord(EncodeEnumDecl(EnumDecl{EdgeID: 6, Field: 5, Code: []Word{1}}))
	if !slices.Equal(ws, []Word{0x0BD5, 0x0051, 6, 0x0051, 5, 1}) {
		t.Fatalf("got %04X", ws)
	}
}
