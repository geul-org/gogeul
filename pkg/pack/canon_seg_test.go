//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestCanonSeg — `0x` segments get A-F upper case (with or without .geul); other segments and `0X` are kept
package pack

import "testing"

func TestCanonSeg(t *testing.T) {
	for in, want := range map[string]string{
		"0x0000abcd":      "0x0000ABCD",
		"0x0000000a.geul": "0x0000000A.geul",
		"0x0000000A.geul": "0x0000000A.geul",
		"versions.geul":   "versions.geul",
		"issuers":         "issuers",
		"0X0000abcd":      "0X0000abcd",
	} {
		if got := canonSeg(in); got != want {
			t.Fatalf("%s: %s", in, got)
		}
	}
}
