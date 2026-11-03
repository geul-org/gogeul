//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestNormalIssuer — 8, 16 and 64 hex digits read as one Issuer at its issued width; other lengths and non-hex fail
package pack

import (
	"strings"
	"testing"
)

func TestNormalIssuer(t *testing.T) {
	for in, want := range map[string]string{
		"0x00000005":                         "0x00000005",
		"0x0000000000000005":                 "0x00000005",
		"0x" + strings.Repeat("0", 63) + "5": "0x00000005",
		"0x0000000100000000":                 "0x0000000100000000",
		"0x000000ab":                         "0x000000AB",
		"0x000000000000aBcD":                 "0x0000ABCD",
	} {
		if got, ok := normalIssuer(in); !ok || got != want {
			t.Fatalf("%s: %s %v", in, got, ok)
		}
	}
	for _, in := range []string{"0x0005", "0x000000000005", "0x0000000g", "0X000000AB", "issuers"} {
		if _, ok := normalIssuer(in); ok {
			t.Fatalf("%s accepted", in)
		}
	}
}
