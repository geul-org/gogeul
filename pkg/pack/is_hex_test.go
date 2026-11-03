//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestIsHex — 0x plus exactly the given number of lowercase hex digits
package pack

import "testing"

func TestIsHex(t *testing.T) {
	for s, want := range map[string]bool{"0x0000abcd": true, "0x0000ABCD": false, "0x000abcd": false, "0000abcd00": false, "0x0000abcg": false} {
		if isHex(s, 8) != want {
			t.Fatalf("%s", s)
		}
	}
}
