//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestNotIdentified — a 5Q code with a quantifier other than Identified; Identified codes, markers and pointers are not
package schema

import "testing"

func TestNotIdentified(t *testing.T) {
	for c, want := range map[uint16]bool{0x05: false, 0x06: true, 0x09: true, 0x1D: true, 0x1E: false, 0x20: false, 0x24: true, 0x39: false, 0x3D: true, 0x3E: false, 0x00: false, 0x04: false, 0x41: false, 0x42: false, 0x5D: false} {
		if notIdentified(c) != want {
			t.Fatalf("C %X", c)
		}
	}
}
