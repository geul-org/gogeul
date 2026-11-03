//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what TestDocIssuer — a document's DOC-ID Issuer; a version list's first PUBLICATION Issuer; none for unreadable bytes, broken documents, empty or foreign bodies
package pack

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocIssuer(t *testing.T) {
	for _, p := range []string{"0x00000000/docs/0x0000000000000845/0x00000001.geul", "0x00000000/docs/0x0000000000000845/versions.geul"} {
		if is, ok := docIssuer(fixtureFile(t, p)); !ok || !slices.Equal(is, []core.Word{0, 0}) {
			t.Fatalf("%s: %v %v", p, is, ok)
		}
	}
	empty, _ := core.EncodeDocument(nil, nil)
	other, _ := core.EncodeDocument(nil, []core.Record{core.EncodeUintValue(1)})
	for i, b := range [][]byte{{1}, {0, 1}, empty, other} {
		if _, ok := docIssuer(b); ok {
			t.Fatalf("bad %d: want none", i)
		}
	}
}
