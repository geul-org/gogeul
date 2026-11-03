//ff:func feature=core type=helper control=iteration dimension=1
//ff:what TestRefSlot — scheme and Quantifier from the section and offset C at the Reference section boundaries (C 0x1D/0x1E/0x1F/0x20, 0x3D/0x3E, 0x40/0x42/0x43, 0x5C/0x5D/0x5F)
package core

import "testing"

func TestRefSlot(t *testing.T) {
	type sq struct {
		s, q uint16
		ok   bool
	}
	for c, want := range map[uint16]sq{
		0x02: {}, 0x03: {}, 0x04: {}, 0x05: {1, 0, true}, 0x1D: {5, 4, true}, 0x1E: {}, 0x1F: {},
		0x20: {6, 0, true}, 0x2E: {8, 4, true}, 0x2F: {9, 0, true}, 0x3D: {11, 4, true}, 0x3E: {}, 0x3F: {},
		0x40: {}, 0x42: {}, 0x43: {}, 0x45: {}, 0x5C: {}, 0x5D: {}, 0x5F: {}, 0x60: {}, 0x65: {},
	} {
		s, q, ok := refSlot(c)
		if (sq{s, q, ok}) != want {
			t.Fatalf("C 0x%02X: %d %d %v, want %+v", c, s, q, ok, want)
		}
	}
}
