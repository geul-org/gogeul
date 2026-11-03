//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDeltaRemove — #7 and #12 are the docs' 1A3B 0088 0007 0088 000C
package core

import (
	"slices"
	"testing"
)

func TestEncodeDeltaRemove(t *testing.T) {
	if ws, _ := EncodeRecord(EncodeDeltaRemove([]uint64{7, 12})); !slices.Equal(ws, []Word{0x1DC4, 0x0051, 7, 0x0051, 12}) {
		t.Fatalf("got %04X", ws)
	}
}
