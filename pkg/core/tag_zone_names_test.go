//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestTagZoneNames — named codes per zone match migration.md §3: Reference 64 with no gap, Value 0x00~0x1D and 0x3D~0x3F, Meta 0x00~0x29, Edge Reserved 0x7C~0x7F and 0x80~0x13C empty
package core

import "testing"

func TestTagZoneNames(t *testing.T) {
	named := func(k Kind, c uint16) bool {
		return CodecName(k, c) != "" || (k == KindReference && SlotMarkerName(c) != "")
	}
	for c := uint16(0); c <= 0x3F; c++ {
		if !named(KindReference, c) {
			t.Fatalf("Reference C 0x%02X unnamed", c)
		}
		if named(KindValue, c) != (c <= 0x21 || c >= 0x3D) {
			t.Fatalf("Value C 0x%02X: %v", c, named(KindValue, c))
		}
		if named(KindMeta, c) != (c <= 0x29) {
			t.Fatalf("Meta C 0x%02X: %v", c, named(KindMeta, c))
		}
	}
	for c := uint16(0x7C); c <= 0x13C; c++ {
		if named(KindEdge, c) {
			t.Fatalf("Edge 0x%03X (Reserved or empty) is named %q", c, CodecName(KindEdge, c))
		}
	}
}
