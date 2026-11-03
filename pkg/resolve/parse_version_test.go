//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestParseVersion — decimal and 2-word hex; other text and an overflow fail
package resolve

import "testing"

func TestParseVersion(t *testing.T) {
	if v, err := parseVersion("12"); err != nil || v != 12 {
		t.Fatal("decimal")
	}
	if v, err := parseVersion("0x0000000C"); err != nil || v != 12 {
		t.Fatal("hex")
	}
	for _, s := range []string{"x", "4294967296", "0xC"} {
		if _, err := parseVersion(s); err == nil {
			t.Fatalf("%s: want error", s)
		}
	}
}
