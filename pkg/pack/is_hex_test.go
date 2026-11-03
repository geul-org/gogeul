//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestIsHex — 0x plus exactly the given number of hex digits, either case
package pack

import "testing"

func TestIsHex(t *testing.T) {
	for s, want := range map[string]bool{"0x0000ABCD": true, "0x0000abcd": true, "0x0000AbCd": true, "0X0000ABCD": false, "0x000ABCD": false, "0000abcd00": false, "0x0000ABCg": false} {
		if isHex(s, 8) != want {
			t.Fatalf("%s", s)
		}
	}
}
