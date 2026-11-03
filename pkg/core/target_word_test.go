//ff:func feature=core type=helper control=sequence
//ff:what TestTargetWord — PUBLICATION is 0x0980, TOMBSTONE 0x09A0, Meta PARENT 0x1D10; SplitTarget inverts it
package core

import "testing"

func TestTargetWord(t *testing.T) {
	if w := TargetWord(KindEdge, EdgePublication); w != 0x0980 {
		t.Fatalf("got %04X", w)
	}
	if w := TargetWord(KindEdge, EdgeTombstone); w != 0x09A0 {
		t.Fatalf("got %04X", w)
	}
	if w := TargetWord(KindMeta, MetaParent); w != 0x1D10 {
		t.Fatalf("got %04X", w)
	}
	if k, c := SplitTarget(TargetWord(KindMeta, 0x30)); k != KindMeta || c != 0x30 {
		t.Fatalf("got %v %02X", k, c)
	}
}
