//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestIsQuantified — Reference C 0x05~0x1D (section 0) and 0x20~0x3D (section 1) are 5Q codes; markers, empty codes and section 2 are not
package core

import "testing"

func TestIsQuantified(t *testing.T) {
	for c := uint16(0); c <= 0x7F; c++ {
		if got := IsQuantified(c); got != (c >= 0x05 && c <= 0x1D || c >= 0x20 && c <= 0x3D) {
			t.Fatalf("C %X: %v", c, got)
		}
	}
}
