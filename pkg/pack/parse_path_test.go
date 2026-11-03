//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestParsePath — the stored 공표 주소 forms map to P, H and named, a zero-padded Issuer to its canonical form; latest.geul, the gone docs/·codecs/ lists and every other spelling are outside
package pack

import (
	"strings"
	"testing"
)

func TestParsePath(t *testing.T) {
	r := parsePath("0x00000001/docs/0x510000000000002A/0x00000001.geul")
	if r.space != spaceP || r.issuer != "0x00000001" || r.doc != 0x510000000000002a || r.slot != 1 {
		t.Fatal(r)
	}
	if r := parsePath("0x0000000000000001/docs/0x510000000000002A/versions.geul"); r.space != spaceP || r.issuer != "0x00000001" || r.path != "0x00000001/docs/0x510000000000002A/versions.geul" {
		t.Fatal(r)
	}
	h := "0x00000000/hashes/0x" + strings.Repeat("0a", 32) + ".geul"
	if r := parsePath(h); r.space != spaceH || len(r.hash) != 32 || r.hash[0] != 0x0a || r.issuer != "0x00000000" {
		t.Fatal(r)
	}
	for in, want := range map[string]string{
		"0x00000000/hashes/0x" + strings.Repeat("0a", 32) + ".geul": "0x00000000/hashes/0x" + strings.Repeat("0A", 32) + ".geul",
		"0x00000000/codecs/0x00000abc/0x0000000a.geul":              "0x00000000/codecs/0x00000ABC/0x0000000A.geul",
		"issuers/0x0000000a.geul":                                   "issuers/0x0000000A.geul",
		"0x000000000000000a/docs/0x510000000000002a/versions.geul":  "0x0000000A/docs/0x510000000000002A/versions.geul",
	} {
		if r := parsePath(in); r.space == spaceNone || r.path != want {
			t.Fatalf("either case %s: %+v", in, r)
		}
	}
	for _, p := range []string{"issuers/0x00000001.geul", "issuers/versions.geul", "0x00000000/codecs/0x00000420/0x00000001.geul", "0x00000000/codecs/0x00000420/versions.geul"} {
		if parsePath(p).space != spaceNamed {
			t.Fatal(p)
		}
	}
	for _, p := range []string{"0x00000000/docs/0x00000001.geul", "0x00000000/docs/versions.geul", "0x00000000/codecs/0x00000002.geul", "0x00000000/codecs/versions.geul",
		"issuers/latest.geul", "0x00000000/docs/latest.geul", "0x00000000/docs/0x0000000000000001/latest.geul",
		"0x00000000/codecs/0x00000420/latest.geul", "issuers.geul", "0x0000/docs/0x00000001.geul", "h/0x" + strings.Repeat("0a", 32) + ".geul",
		"0X00000000/hashes/0x" + strings.Repeat("0A", 32) + ".geul", "0x00000000/hashes/0x" + strings.Repeat("0a", 32) + ".gel",
		"0xZZ/docs/0x0000000000000001/versions.geul", "0x00000000/0x0000000000000001/versions.geul", "0x00000000/keys/x.geul",
		"0x00000000/docs/0x00000000.geul", "0x00000000/docs/0xFFFFFFFFFFFFFFFF/versions.geul", "0x00000000/docs/0xffffffffffffffff/versions.geul", "0x00000001/docs/0xFFFFFFFFFFFFFFFF/0x00000001.geul", "0x00000000/codecs/0x420/0x00000001.geul", "issuers/0x00000001.geul/x", "latest.txt"} {
		if parsePath(p).space != spaceNone {
			t.Fatal(p)
		}
	}
}
