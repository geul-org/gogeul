//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestParseHPath — a 64-digit hash file name is the H space with the 32-byte hash; other names leave no space
package pack

import (
	"strings"
	"testing"
)

func TestParseHPath(t *testing.T) {
	h := "0x" + strings.Repeat("ab", 32)
	ref := parseHPath(pathRef{issuer: "0x00000000"}, h+".geul")
	if ref.space != spaceH || ref.issuer != "0x00000000" || ref.hash != strings.Repeat("\xab", 32) {
		t.Fatalf("%+v", ref)
	}
	for _, n := range []string{h, h[:64] + ".geul", strings.ToUpper(h) + ".geul", "latest.geul"} {
		if r := parseHPath(pathRef{issuer: "0x00000000"}, n); r.space != spaceNone {
			t.Fatalf("%s: %+v", n, r)
		}
	}
}
