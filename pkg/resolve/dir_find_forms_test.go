//ff:func feature=resolve type=resolver control=iteration dimension=2
//ff:what TestDirFindForms — every 공표 주소 gives the file tree's bytes from every storage form; a missing document is ErrNotFound in all of them
package resolve

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirFindForms(t *testing.T) {
	tree := formsTree(t)
	const id, gone, other = "0x00000000/docs/0xffff000000000007", "0x00000000/docs/0xffff000000000008", "0x00000002/docs/0x0000000000000005"
	v1, err := tree.Find(id + "/0x00000001.geul")
	if err != nil {
		t.Fatal(err)
	}
	ws, _ := core.ReadWords(v1)
	sum, _ := core.ContentHash(ws)
	od, err := tree.Find(other + "/latest.geul")
	if err != nil {
		t.Fatal(err)
	}
	ows, _ := core.ReadWords(od)
	osum, _ := core.ContentHash(ows)
	addrs := []string{"issuers/latest.geul", "/.well-known/geul/issuers/0x00000001.geul", "0x00000000/docs/latest.geul", "geul.org/docs/latest.geul",
		id + "/latest.geul", "geul.org/docs/seven/latest.geul", id + "/0x00000001.geul", id + "/0x00000002", "geul.org/docs/seven/0x00000001.geul",
		id + "/versions.geul", "geul.org/docs/seven/versions.geul", gone + "/versions.geul",
		other + "/latest.geul", other + "/0x00000001.geul", other + "/versions.geul", "other.example/docs/0x0000000000000005/latest.geul",
		hashPath(geulIssuer, sum), hashPath("0x00000002", osum)}
	missing := []string{id + "/0x00000009.geul", gone + "/0x00000001.geul", gone + "/latest.geul", "0x00000000/docs/0xffff000000000009/latest.geul",
		"0x00000009/docs/0x0000000000000001/versions.geul", "0x00000009/docs/latest.geul", "geul.org/docs/nothing/latest.geul", "nobody.example/docs/seven/latest.geul",
		"h/0x" + strings.Repeat("0", 64) + ".geul", "issuers.geul", "0x00000000/documents.geul", hashPath("0x00000002", sum)}
	for _, f := range formsOf(t, tree) {
		for _, a := range addrs {
			want, err := tree.Find(a)
			if err != nil {
				t.Fatal(a, err)
			}
			if b, err := f.dir.Find(a); err != nil || !bytes.Equal(b, want) {
				t.Fatalf("%s: %s: %d bytes, %v", f.name, a, len(b), err)
			}
		}
		for _, a := range missing {
			if _, err := f.dir.Find(a); !errors.Is(err, ErrNotFound) {
				t.Fatalf("%s: %s: %v", f.name, a, err)
			}
		}
	}
}
