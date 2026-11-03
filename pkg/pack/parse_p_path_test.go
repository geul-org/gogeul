//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestParsePPath — versions.geul is slot 0, a non-zero 8-digit version is its slot; other spellings, latest.geul included, are outside the forms
package pack

import "testing"

func TestParsePPath(t *testing.T) {
	ref := pathRef{issuer: "0x00000000"}
	if r := parsePPath(ref, "0x0000000000000008", "versions.geul"); r.space != spaceP || r.slot != 0 || r.doc != 8 {
		t.Fatal(r)
	}
	if r := parsePPath(ref, "0x0000000000000008", "0x0000000a.geul"); r.space != spaceP || r.slot != 10 {
		t.Fatal(r)
	}
	for _, c := range [][2]string{{"0x08", "versions.geul"}, {"0x0000000000000008", "0x00000000.geul"},
		{"0x0000000000000008", "0x00000001.txt"}, {"0x0000000000000008", "x.geul"}, {"0x0000000000000008", "latest.geul"}} {
		if r := parsePPath(ref, c[0], c[1]); r.space != spaceNone {
			t.Fatal(c)
		}
	}
}
