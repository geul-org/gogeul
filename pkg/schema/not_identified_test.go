//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestNotIdentified — a 5Q code with a quantifier other than Identified; Identified codes, markers and pointers are not
package schema

import "testing"

func TestNotIdentified(t *testing.T) {
	for c, want := range map[uint16]bool{0x05: false, 0x06: true, 0x09: true, 0x37: false, 0x3B: true, 0x00: false, 0x04: false, 0x3C: false} {
		if notIdentified(c) != want {
			t.Fatalf("C %X", c)
		}
	}
}
