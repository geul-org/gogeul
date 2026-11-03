//ff:func feature=pack type=writer control=sequence
//ff:what TestSplitHome — 0x0000 first, else the first Issuer in Dir order, else the top set
package pack

import "testing"

func TestSplitHome(t *testing.T) {
	if splitHome(map[string]bool{"0x00000001": true, "0x00000000": true}) != "0x00000000" || splitHome(map[string]bool{"0x00000001": true, "0x00000002": true}) != "0x00000001" || splitHome(nil) != "" {
		t.Fatal("splitHome")
	}
}
