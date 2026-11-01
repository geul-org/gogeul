//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirFind — every 공표 주소 form finds its document by ID or name; a malformed address, version or hash and a missing document fail
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
	issuers, _ := d.Fetch(t.Context(), issuersFile)
	names, _ := d.Fetch(t.Context(), "0x0000/"+namesFile)
	list, _ := d.Fetch(t.Context(), listPath(sampleKeyOf(7)))
	ws, _ := core.ReadWords(v1)
	sum, _ := core.ContentHash(ws)
	h := hashPath(sum)
	const id = "0x0000/0xffff000000000007"
	for addr, want := range map[string][]byte{
		"issuers.geul":                      issuers,
		"/.well-known/geul/issuers.geul":    issuers,
		"0x0000/documents.geul":             names,
		"geul.org/documents.geul":           names,
		id + ".geul":                        v2,
		id:                                  v2,
		"geul.org/seven.geul":               v2,
		"/.well-known/geul/" + id + ".geul": v2,
		id + "/0x00000001.geul":             v1,
		id + "/0x00000001":                  v1,
		"geul.org/seven/0x00000001.geul":    v1,
		id + "/0x00000002.geul":             v2,
		id + "/0x00000000.geul":             v2,
		id + "/versions.geul":               list,
		"geul.org/seven/versions.geul":      list,
		h:                                   v1,
	} {
		b, err := d.Find(addr)
		if err != nil || !bytes.Equal(b, want) {
			t.Fatalf("%s: %d bytes, %v", addr, len(b), err)
		}
	}
	for _, addr := range []string{"seven", "", "h/0x12", id + "/1", id + "@1", id + "/0x00000001/x", "0x0000/0xFFFF000000000007", "nobody/seven", id + "/"} {
		if _, err := d.Find(addr); err == nil {
			t.Fatalf("%q: want error", addr)
		}
	}
	for _, addr := range []string{id + "/0x00000009.geul", "0x0000/0xffff000000000009.geul", "0x0001/documents.geul", "geul.org/nothing", "h/0x" + string(bytes.Repeat([]byte("0"), 64)) + ".geul"} {
		if _, err := d.Find(addr); !errors.Is(err, ErrNotFound) {
			t.Fatalf("%q: %v", addr, err)
		}
	}
}
