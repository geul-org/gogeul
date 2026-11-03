//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeNextLocalID — 524 is the docs' 18BC 020C; 70000 takes two words
package core

import (
	"slices"
	"testing"
)

func TestEncodeNextLocalID(t *testing.T) {
	if ws, _ := EncodeRecord(EncodeNextLocalID(524)); !slices.Equal(ws, []Word{0x1DD1, 0x020C}) {
		t.Fatalf("524: %04X", ws)
	}
	if r := EncodeNextLocalID(70000); len(r.Payload) != 2 {
		t.Fatalf("70000: %04X", r.Payload)
	}
}
