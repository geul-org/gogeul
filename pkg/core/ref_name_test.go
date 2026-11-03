//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefName — a 5Q code is named by its scheme (any quantifier, either section), the section 2 codes by their own name, empty codes ""
package core

import "testing"

func TestRefName(t *testing.T) {
	for c, want := range map[uint16]string{
		0x05: "LOCAL", 0x09: "LOCAL", 0x1D: "WIKIDATA-PID", 0x20: "EXTERNAL-POINTER-1WORD", 0x24: "EXTERNAL-POINTER-1WORD",
		0x25: "EXTERNAL-POINTER-2WORD", 0x39: "EXTERNAL-ISSUER-16WORD", 0x3D: "EXTERNAL-ISSUER-16WORD",
		0x40: "ISSUER-POINTER", 0x41: "DOC-POINTER", 0x42: "CODEC-POINTER", 0x5D: "REFERENCE-ISSUER-2WORD",
		0x03: "", 0x1E: "", 0x3E: "", 0x43: "", 0x5C: "",
	} {
		if got := refName(c); got != want {
			t.Fatalf("C %X: %q, want %q", c, got, want)
		}
	}
}
