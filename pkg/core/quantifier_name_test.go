//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestQuantifierName — 5Q names by C % 5 (C the offset in the section) in sections 0 and 1, "" for markers, empty codes, pointers, CODEC-POINTER and REFERENCE-ISSUER
package core

import "testing"

func TestQuantifierName(t *testing.T) {
	cases := map[uint16]string{
		0: "", 1: "", 2: "", 3: "", 4: "",
		0x05: "Identified", 0x06: "Generic", 0x07: "Free-choice", 0x08: "Universal", 0x09: "Variable",
		0x14: "Identified", 0x18: "Variable", 0x19: "Identified", 0x1D: "Variable", 0x1E: "", 0x1F: "",
		0x20: "Identified", 0x21: "Generic", 0x24: "Variable", 0x25: "Identified", 0x2E: "Variable", 0x2F: "Identified",
		0x38: "Variable", 0x39: "Identified", 0x3C: "Universal", 0x3D: "Variable", 0x3E: "", 0x3F: "",
		0x40: "", 0x41: "", 0x42: "", 0x43: "", 0x45: "", 0x5C: "", 0x5D: "", 0x5E: "", 0x5F: "",
	}
	for c, want := range cases {
		if got := QuantifierName(c); got != want {
			t.Fatalf("QuantifierName(0x%02X)=%q want %q", c, got, want)
		}
	}
}
