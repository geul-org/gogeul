//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestQuantifierName — 5Q names by C%5 for 5 <= C < 0x3C, "" for markers, pointers, CODEC-POINTER and REFERENCE-ISSUER
package core

import "testing"

func TestQuantifierName(t *testing.T) {
	cases := map[uint16]string{
		0: "", 1: "", 2: "", 3: "", 4: "",
		0x05: "Identified", 0x06: "Generic", 0x07: "Free-choice", 0x08: "Universal", 0x09: "Variable",
		0x14: "Identified", 0x18: "Variable", 0x1D: "Variable", 0x1E: "Identified", 0x2C: "Variable", 0x2D: "Identified",
		0x36: "Variable", 0x37: "Identified", 0x3B: "Variable",
		0x3C: "", 0x3D: "", 0x3E: "", 0x3F: "",
	}
	for c, want := range cases {
		if got := QuantifierName(c); got != want {
			t.Fatalf("QuantifierName(0x%02X)=%q want %q", c, got, want)
		}
	}
}
