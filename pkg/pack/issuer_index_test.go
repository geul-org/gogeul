//ff:func feature=pack type=helper control=sequence
//ff:what TestIssuerIndex — position by ID
package pack

import "testing"

func TestIssuerIndex(t *testing.T) {
	m := issuerIndex([]string{"0x00000000", "0x00000001"})
	if m["0x00000001"] != 1 || len(m) != 2 {
		t.Fatal(m)
	}
}
