//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsQuantified — Reference C 0x05~0x3B are 5Q codes; markers, pointers, CODEC-POINTER and REFERENCE-ISSUER are not
package core

import "testing"

func TestIsQuantified(t *testing.T) {
	for c := uint16(0); c <= 0x3F; c++ {
		if got := IsQuantified(c); got != (c >= 0x05 && c <= 0x3B) {
			t.Fatalf("C %X: %v", c, got)
		}
	}
}
