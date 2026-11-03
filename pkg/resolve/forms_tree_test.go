//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what formsTree — test fixture: two Issuers named in the Issuer list, a two-version document, a deleted version and long compressible texts, as a file tree
package resolve

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// formsTree publishes the documents every storage form must serve alike.
func formsTree(t *testing.T) Dir {
	t.Helper()
	d := Dir{Root: t.TempDir()}
	s0, _ := core.EncodeIssuerPointer([]core.Word{0, 0})
	s1, _ := core.EncodeIssuerPointer([]core.Word{0, 2})
	b := bindDoc(core.Bind{Subject: s0, Name: []byte("geul.org")}, core.Bind{Subject: s1, Name: []byte("other.example")})
	if _, err := d.PublishList([]core.Word{0, 0}, core.DocIDIssuers, b); err != nil {
		t.Fatal(err)
	}
	long := strings.Repeat("글은 언어가 아니라 글자다 ", 200)
	for _, b := range [][]byte{edgeDoc(7, long), edgeDoc(7, long, "b"), edgeDoc(8, "gone"), otherDoc(long)} {
		if _, err := d.Publish(b); err != nil {
			t.Fatal(err)
		}
	}
	if err := d.Tombstone([]core.Word{0, 0}, []core.Word{0xffff, 0, 0, 8}, 1, 1); err != nil {
		t.Fatal(err)
	}
	return d
}
