//ff:func feature=server type=helper control=iteration dimension=1
//ff:what TestLatestDocs — the newest version of an Issuer's document list through its versions.geul; a name and a missing list fail
package server

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestLatestDocs(t *testing.T) {
	d := resolve.Dir{Root: t.TempDir()}
	v1, _ := core.EncodeDocument(nil, nil)
	v2, _ := core.EncodeDocument(nil, []core.Record{core.EncodeUintValue(2)})
	for _, b := range [][]byte{v1, v2} {
		if _, err := d.PublishList([]core.Word{0, 2}, core.DocIDDocs, b); err != nil {
			t.Fatal(err)
		}
	}
	if b, err := latestDocs(d.Tree(), "0x00000002"); err != nil || !bytes.Equal(b, v2) {
		t.Fatal(len(b), err)
	}
	if _, err := latestDocs(d.Tree(), "geul.org"); err == nil {
		t.Fatal("name: want error")
	}
	if _, err := latestDocs(d.Tree(), "0x00000003"); err == nil {
		t.Fatal("missing: want error")
	}
}
