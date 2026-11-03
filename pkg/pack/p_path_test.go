//ff:func feature=pack type=encoder control=sequence
//ff:what TestPPath — slot 0 is versions.geul, others 8-digit version files
package pack

import "testing"

func TestPPath(t *testing.T) {
	if p := pPath("0x00000000", 8, 0); p != "0x00000000/docs/0x0000000000000008/versions.geul" {
		t.Fatal(p)
	}
	if p := pPath("0x00000000", 8, 2); p != "0x00000000/docs/0x0000000000000008/0x00000002.geul" {
		t.Fatal(p)
	}
}
