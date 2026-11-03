//ff:func feature=pack type=decoder control=iteration dimension=1
//ff:what TestParseHex — 0x hex of the exact width only, digits in either case
package pack

import "testing"

func TestParseHex(t *testing.T) {
	if b, ok := parseHex("0x00FF", 4); !ok || b[1] != 0xff {
		t.Fatal("valid")
	}
	if b, ok := parseHex("0x00fF", 4); !ok || b[1] != 0xff {
		t.Fatal("either case")
	}
	for _, s := range []string{"0x0FF", "0X00ff", "0X00FF", "1x00ff", "0x00Fg"} {
		if _, ok := parseHex(s, 4); ok {
			t.Fatal(s)
		}
	}
}
