//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestGlobeCodec — exactly the ten codes of COORDINATE-GEO's globe list among all 96 Reference codes
package schema

import "testing"

func TestGlobeCodec(t *testing.T) {
	want := map[uint16]bool{0x05: true, 0x0A: true, 0x0F: true, 0x14: true, 0x20: true, 0x25: true, 0x2A: true, 0x2F: true, 0x34: true, 0x39: true}
	for c := uint16(0); c < 0x60; c++ {
		if globeCodec(c) != want[c] {
			t.Fatalf("C 0x%02X: %v", c, globeCodec(c))
		}
	}
}
