//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestCanonicalIssuer — only the issued-width lowercase spelling is canonical; zero-padded wider spellings and other text are not
package pack

import (
	"strings"
	"testing"
)

func TestCanonicalIssuer(t *testing.T) {
	for id, want := range map[string]bool{
		"0x00000000": true, "0x0000000100000000": true, "0x1" + strings.Repeat("0", 63): true,
		"0x0000000000000000": false, "0x00000000000000000000000000000000000000000000000000000000000000ab": false, "0x000000AB": false, "issuers": false,
	} {
		if got := canonicalIssuer(id); got != want {
			t.Fatalf("%s: %v", id, got)
		}
	}
}
