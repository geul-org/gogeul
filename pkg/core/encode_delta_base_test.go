//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDeltaBase — base version 3 is the docs' 193A 0000 0003
package core

import (
	"slices"
	"testing"
)

func TestEncodeDeltaBase(t *testing.T) {
	if ws, _ := EncodeRecord(EncodeDeltaBase(3)); !slices.Equal(ws, []Word{0x1DB2, 0, 3}) {
		t.Fatalf("got %04X", ws)
	}
}
