//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestTagVector — every new TAG of the reference vector decodes to the Kind, name and L of the table and re-encodes to itself; new TAGs are unique
package core

import "testing"

func TestTagVector(t *testing.T) {
	rows := loadVector(t)
	if len(rows) != 3315 {
		t.Fatalf("vector rows %d, want 3315 (old valid TAGs)", len(rows))
	}
	seen := map[int]Word{}
	for _, r := range rows {
		if r.New < 0 {
			continue
		}
		if prev, dup := seen[r.New]; dup {
			t.Fatalf("new %04X from both %04X and %04X", r.New, prev, r.Old)
		}
		seen[r.New] = r.Old
		w := Word(r.New)
		k, name, q, l := vectorExpect(r.Name, r.OldKey)
		gk, codec := SplitTarget(w)
		if gk != k || KindOf(w) != k || int(w&0xF) != l || tagWord(k, uint8(l), codec) != w {
			t.Fatalf("%04X (%s): kind %v codec %03X", w, r.Name, gk, codec)
		}
		got := CodecName(k, codec)
		if k == KindReference && IsSlotMarkerCodec(codec) {
			got = SlotMarkerName(codec)
		}
		if got != name && !(k == KindEdge && got == "" && coreUnnamedEdge(codec)) {
			t.Fatalf("%04X: name %q want %q", w, got, name)
		}
		if QuantifierName(codec) != q && k == KindReference {
			t.Fatalf("%04X: quantifier %q want %q", w, QuantifierName(codec), q)
		}
		vectorRoundTrip(t, w, k, codec, l)
	}
	if len(seen) != 3315-192 {
		t.Fatalf("mapped %d, want %d", len(seen), 3315-192)
	}
}
