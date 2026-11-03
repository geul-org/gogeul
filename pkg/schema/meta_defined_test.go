//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestMetaDefined — Meta C 0x00~0x29 are assigned, 0x2A~0x3F are not
package schema

import "testing"

func TestMetaDefined(t *testing.T) {
	for c := uint16(0); c <= 0x3F; c++ {
		if metaDefined(c) != (c <= 0x29) {
			t.Fatalf("0x%02X: %v", c, metaDefined(c))
		}
	}
}
