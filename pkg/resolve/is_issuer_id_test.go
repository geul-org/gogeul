//ff:func feature=resolve type=decoder control=iteration dimension=1
//ff:what TestIsIssuerID — 1, 2, 4 and 16-word IDs pass; 3 words and bad hex fail
package resolve

import (
	"strings"
	"testing"
)

func TestIsIssuerID(t *testing.T) {
	for s, want := range map[string]bool{
		"0x00000002":                   true,
		"0x00000001":                   true,
		"0x0000000000000001":           true,
		"0x" + strings.Repeat("a", 64): true,
		"0x000000000001":               false,
		"0x000g":                       false,
	} {
		if isIssuerID(s) != want {
			t.Fatalf("%q: want %v", s, want)
		}
	}
}
