//ff:func feature=core type=encoder control=sequence
//ff:what TestEncodeDeclare — the PUBLICATION example words of edge/DECLARE: Edge Id and target only
package core

import (
	"slices"
	"testing"
)

func TestEncodeDeclare(t *testing.T) {
	ws, _ := EncodeRecord(EncodeDeclare(Declare{EdgeID: 1, Target: 0x0980}))
	want := []Word{0x09B3, 0x0051, 1, 0x0980}
	if !slices.Equal(ws, want) {
		t.Fatalf("got %04X", ws)
	}
}
