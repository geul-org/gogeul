//ff:func feature=pack type=helper control=sequence
//ff:what TestIssuerLess — width first, then hex
package pack

import "testing"

func TestIssuerLess(t *testing.T) {
	if !issuerLess("0xFFFF", "0x00000000") || !issuerLess("0x00000000", "0x00000002") || issuerLess("0x00000002", "0x00000000") {
		t.Fatal("issuerLess")
	}
}
