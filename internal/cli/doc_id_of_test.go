//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestDocIDOf — the DOC-ID's Issuer and document; odd bytes, a broken document and none fail
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDocIDOf(t *testing.T) {
	id, _ := core.EncodeDocID([]core.Word{0, 0}, []core.Word{0, 0, 0, 7})
	b, _ := core.EncodeDocument([]core.Record{core.EncodeSpecVersion(2), id}, nil)
	if is, doc, err := docIDOf(b); err != nil || is[0] != 0 || doc[3] != 7 {
		t.Fatal(err)
	}
	none, _ := core.EncodeDocument(nil, nil)
	for _, x := range [][]byte{{1}, {0, 5}, none} {
		if _, _, err := docIDOf(x); err == nil {
			t.Errorf("%v: want error", x)
		}
	}
}
