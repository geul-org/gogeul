//ff:func feature=pack type=writer control=sequence
//ff:what TestSplitPaths — Issuer paths (hashes/ included) go to their Issuer, a zero-padded Issuer to the same set, the Issuer list to the home set
package pack

import "testing"

func TestSplitPaths(t *testing.T) {
	g := splitPaths([]string{"0x00000000/docs/0x0000000000000001/versions.geul", "0x0000000000000000/hashes/0x" + "ab" + "00000000000000000000000000000000000000000000000000000000000000.geul",
		"0x00000001/docs/0x0000000000000002/0x00000001.geul", "issuers/0x00000001.geul"})
	if len(g) != 2 || len(g["0x00000000"]) != 3 || len(g["0x00000001"]) != 1 {
		t.Fatal(g)
	}
	if top := splitPaths([]string{"issuers/0x00000001.geul"}); len(top[""]) != 1 {
		t.Fatal(top)
	}
}
