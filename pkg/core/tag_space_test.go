//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestTagSpace — over every TAG 0x0000~0x1FFF: Kind by section, TAG re-packing, ISSUER widths, quantified Reference codes and Edge family follow the plan2 layout table (layout2.md); validation Meta is C >= 0x25
package core

import (
	"strings"
	"testing"
)

func TestTagSpace(t *testing.T) {
	l := loadLayout(t)
	width := map[string]int{"2WORD": 2, "4WORD": 4, "16WORD": 16}
	for v := 0; v < 0x2000; v++ {
		w := Word(v)
		name := l.names[uint16(v>>4)]
		k, c := SplitTarget(w)
		if k != l.kind[v>>9] || KindOf(w) != k || tagWord(k, uint8(w&0xF), c) != w || c > CodecMax(k) {
			t.Fatalf("%04X: %v %03X", v, k, c)
		}
		iw := 0
		if i := strings.LastIndex(name, "-ISSUER-"); i >= 0 && !strings.HasPrefix(name, "EXTERNAL-") {
			iw = width[name[i+len("-ISSUER-"):]]
		}
		if IssuerWidth(k, c) != iw {
			t.Fatalf("%04X %s: IssuerWidth %d want %d", v, name, IssuerWidth(k, c), iw)
		}
		if IsValidationMeta(k, c) != (v >= 0x1E50) {
			t.Fatalf("%04X: validation %v", v, IsValidationMeta(k, c))
		}
		if k == KindEdge && int(c>>5) != v>>9-5 {
			t.Fatalf("%04X: Edge family %d want section %d - 5", v, c>>5, v>>9)
		}
		if k == KindReference && IsQuantified(c) != strings.Contains(name, "×") {
			t.Fatalf("%04X %q: quantified %v", v, name, IsQuantified(c))
		}
	}
}
