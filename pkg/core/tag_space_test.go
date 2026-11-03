//ff:func feature=core type=decoder control=iteration dimension=1
//ff:what TestTagSpace — over every TAG 0x0000~0x1FFF: Kind by zone, TAG re-packing, ISSUER widths, validation Meta, Edge family bits and the quantified Reference range follow migration.md §2.1
package core

import "testing"

func TestTagSpace(t *testing.T) {
	zone := func(w Word) Kind {
		switch {
		case w < 0x0400:
			return KindReference
		case w < 0x0800:
			return KindValue
		case w < 0x1C00:
			return KindEdge
		default:
			return KindMeta
		}
	}
	for v := 0; v < 0x2000; v++ {
		w := Word(v)
		k, c := SplitTarget(w)
		if k != zone(w) || KindOf(w) != k || tagWord(k, uint8(w&0xF), c) != w || c > CodecMax(k) {
			t.Fatalf("%04X: %v %03X", v, k, c)
		}
		iw := 0
		if (v >= 0x03D0 && v < 0x0400) || (v >= 0x07D0 && v < 0x0800) || (v >= 0x1BD0 && v < 0x1C00) {
			iw = [4]int{0, 2, 4, 16}[v>>4&3]
		}
		if IssuerWidth(k, c) != iw {
			t.Fatalf("%04X: IssuerWidth %d want %d", v, IssuerWidth(k, c), iw)
		}
		if IsValidationMeta(k, c) != (v >= 0x1E50) {
			t.Fatalf("%04X: validation %v", v, IsValidationMeta(k, c))
		}
		if v >= 0x0800 && v < 0x1000 && Word(v>>9&3) != Word(c>>5) {
			t.Fatalf("%04X: Edge family %d want %d", v, c>>5, v>>9&3)
		}
		if k == KindReference && IsQuantified(c) != (v >= 0x0050 && v < 0x03C0) {
			t.Fatalf("%04X: quantified %v", v, IsQuantified(c))
		}
	}
}
