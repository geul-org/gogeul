//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirFind — every 공표 주소 form reads its document (names for the Issuer and the document, .geul optional, padded Issuer); other spellings and missing addresses fail
package resolve

import (
	"bytes"
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirFind(t *testing.T) {
	d := issuersFixture(t)
	v1, v2 := edgeDoc(7, "a"), edgeDoc(7, "a", "b")
	d.Publish(v2)
	issuers, _ := d.Fetch(t.Context(), issuersPath())
	names, _ := d.Fetch(t.Context(), namesPath(geulIssuer))
	list, _ := d.Fetch(t.Context(), listPath(sampleKeyOf(7)))
	ws, _ := core.ReadWords(v1)
	sum, _ := core.ContentHash(ws)
	h := hashPath(geulIssuer, sum)
	const id = "0x00000000/docs/0xffff000000000007"
	for addr, want := range map[string][]byte{
		"issuers/latest.geul":                                    issuers,
		"/.well-known/geul/issuers/0x00000001.geul":              issuers,
		"0x00000000/docs/latest.geul":                            names,
		"geul.org/docs/latest":                                   names,
		id + "/latest.geul":                                      v2,
		"geul.org/docs/seven/latest.geul":                        v2,
		"0x0000000000000000/docs/0xffff000000000007/latest.geul": v2,
		"/.well-known/geul/" + id + "/latest.geul":               v2,
		id + "/0x00000001.geul":                                  v1,
		id + "/0x00000001":                                       v1,
		"geul.org/docs/seven/0x00000001.geul":                    v1,
		id + "/0x00000002.geul":                                  v2,
		id + "/versions.geul":                                    list,
		"geul.org/docs/seven/versions.geul":                      list,
		h:                                                        v1,
	} {
		b, err := d.Find(addr)
		if err != nil || !bytes.Equal(b, want) {
			t.Fatalf("%s: %d bytes, %v", addr, len(b), err)
		}
	}
	for _, addr := range []string{"seven", "", "issuers.geul", "h/0x12", id + "/1", id + "@1", id + "/0x00000001/x", id + "/0x00000000.geul",
		"0x00000000/0xffff000000000007/0x00000001.geul", "0x00000000/documents.geul", "nobody/docs/seven/latest.geul", id + "/",
		id + "/0x00000009.geul", "0x00000000/docs/0xffff000000000009/latest.geul", "0x00000002/docs/latest.geul",
		"geul.org/docs/nothing/latest.geul", "0x00000000/hashes/0x" + string(bytes.Repeat([]byte("0"), 64)) + ".geul"} {
		if _, err := d.Find(addr); !errors.Is(err, ErrNotFound) {
			t.Fatalf("%q: %v", addr, err)
		}
	}
}
