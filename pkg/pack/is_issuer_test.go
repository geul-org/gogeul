//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestIsIssuer — 8, 16 and 64 lowercase hex digits are Issuers (zero padding included); 4 digits and other spellings are not
package pack

import (
	"strings"
	"testing"
)

func TestIsIssuer(t *testing.T) {
	for _, s := range []string{"0x00000000", "0x00000001", "0x0000000000000001", "0x" + strings.Repeat("ab", 32), "0x" + strings.Repeat("0", 63) + "1"} {
		if !isIssuer(s) {
			t.Fatal(s)
		}
	}
	for _, s := range []string{"0x0000", "0x000", "0x000000000000", "0x000000AB", "h", "issuers"} {
		if isIssuer(s) {
			t.Fatal(s)
		}
	}
	if n, _ := normalIssuer("0x" + strings.Repeat("0", 63) + "1"); n != "0x00000001" {
		t.Fatal(n)
	}
}
