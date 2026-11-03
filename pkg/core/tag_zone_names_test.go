//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestTagZoneNames — every 9-bit code 0x000~0x1FF is named as the plan2 name table (layout2.md) says: scheme and Quantifier for 5Q codes, "" for empty, Reserved and the TOULMIN sections (pkg/toulmin names those)
package core

import (
	"strings"
	"testing"
)

func TestTagZoneNames(t *testing.T) {
	l := loadLayout(t)
	named := 0
	for code := uint16(0); code < 0x200; code++ {
		k, c := SplitTarget(Word(code << 4))
		want, q, _ := strings.Cut(l.names[code], "×")
		for _, p := range []string{"V:", "E:", "M:"} {
			want = strings.TrimPrefix(want, p)
		}
		want = strings.TrimSuffix(strings.TrimSuffix(strings.TrimSuffix(want, "-EVEN"), "-ODD"), "Reserved")
		if k == KindEdge && strings.Contains(l.title[code>>5], "TOULMIN") {
			want = ""
		}
		got := CodecName(k, c)
		if k == KindReference && IsSlotMarkerCodec(c) {
			got = SlotMarkerName(c)
		}
		if got != want || k == KindReference && QuantifierName(c) != q {
			t.Fatalf("0x%03XL: %q %q, want %q %q", code, got, QuantifierName(c), want, q)
		}
		if got != "" {
			named++
		}
	}
	if named != 231-43 {
		t.Fatalf("%d named codes, want 188 (231 less the 43 codes of the TOULMIN sections, 4 Reserved among them)", named)
	}
}
