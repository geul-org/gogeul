//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestParseAddr — the ten 공표 주소 forms parse and Path writes them back; 8/16/64-digit Issuers are one Issuer; the gone docs/·codecs/ lists and every other spelling are ErrNotFound
package resolve

import (
	"errors"
	"strings"
	"testing"
)

func TestParseAddr(t *testing.T) {
	h := "0x" + strings.Repeat("ab", 32)
	forms := []string{
		"issuers/latest.geul", "issuers/0x00000002.geul", "issuers/versions.geul",
		"0x00000000/codecs/0x000003c0/latest.geul", "0x00000000/codecs/0x000003c0/0x00000003.geul", "0x00000000/codecs/0x000003c0/versions.geul",
		"0x00000001/docs/0x510000000000002a/latest.geul", "0x00000001/docs/0x510000000000002a/0x00000001.geul", "0x00000001/docs/0x510000000000002a/versions.geul",
		"0x00000001/hashes/" + h + ".geul", "0x00000002/docs/0xfffffffffffffffd/0x00000001.geul",
	}
	var spelled, want []string
	for _, f := range forms {
		spelled = append(spelled, f, ".well-known/geul/"+f, "file:///srv/.well-known/geul/"+f, "https://geul.org/.well-known/geul/"+f)
		want = append(want, f, f, f, f)
	}
	for i, s := range spelled {
		if a, err := ParseAddr(s); err != nil || a.Path() != want[i] {
			t.Fatalf("%s: %+v %v", s, a, err)
		}
	}
	if a, _ := ParseAddr("https://geul.org/.well-known/geul/issuers/latest.geul"); a.Base != "https://geul.org" || a.Kind != AddrIssuers || a.Version != Latest {
		t.Fatalf("base: %+v", a)
	}
	for _, is := range []string{"0x0000000000000001", "0x" + strings.Repeat("0", 63) + "1"} {
		if a, err := ParseAddr(is + "/docs/0x510000000000002a/0x00000001.geul"); err != nil || a.Path() != forms[7] {
			t.Fatalf("%s: %+v %v", is, a, err)
		}
	}
	if a, _ := ParseAddr("0x00000001/docs/0x0000000000000001/versions.geul"); a.Kind != AddrDoc || !a.Versions || a.ID[3] != 1 {
		t.Fatalf("Issuer and document 1 apart: %+v", a)
	}
	for _, s := range []string{"", "issuers.geul", "0x0000/documents.geul", "h/" + h + ".geul", "0x0000/docs/latest.geul", "0x00000000/docs/0x00000000.geul",
		"0x00000000/docs/0x0000000000000001/0x00000000.geul", "0x00000000/0x0000000000000001/0x00000001.geul", "0x00000000/docs/0x0000000000000001.geul",
		"0x00000000/codecs/0x3c0/latest.geul", "0x00000000/codecs/0x000003c0/0x000003c0/latest.geul", "0x00000000/keys/x.geul", "0x00000000/hashes/0xab.geul",
		"0x00000000/docs/latest", "0x00000000/DOCS/latest.geul", "0x0000000A/docs/latest.geul", "0x00000000/docs/0x00000001.geul/x",
		"0x00000000/codecs/latest.geul", "0x00000000/codecs/0x00000001.geul", "0x00000000/codecs/versions.geul",
		"0x00000001/docs/latest.geul", "0x00000001/docs/0x00000001.geul", "0x00000001/docs/versions.geul",
		"geul.org/docs/latest.geul", "0x00000000/docs/0xffffffffffffffff/latest.geul", "http://geul.org/.well-known/geul/issuers/latest.geul", "https://geul.org/issuers/latest.geul", "issuers/0x00000002.GEUL"} {
		if _, err := ParseAddr(s); !errors.Is(err, ErrNotFound) {
			t.Fatalf("%q: %v", s, err)
		}
	}
}
