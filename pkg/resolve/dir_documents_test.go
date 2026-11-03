//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirDocuments — document directories with names; files and other entries skipped; a missing Issuer and a broken index fail
package resolve

import (
	"os"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirDocuments(t *testing.T) {
	d := issuersFixture(t)
	os.MkdirAll(d.file("0x00000000/docs/x"), 0o755)
	got, err := d.Documents([]core.Word{0, 0})
	if err != nil || len(got) != 1 || got[0].ID != "0xffff000000000007" || got[0].Names[0] != "seven" {
		t.Fatalf("%+v %v", got, err)
	}
	if _, err := d.Documents([]core.Word{0, 1}); err == nil {
		t.Fatal("missing Issuer: want error")
	}
	d.write(listPath(namesKey(geulIssuer)), []byte{1})
	if _, err := d.Documents([]core.Word{0, 0}); err == nil {
		t.Fatal("broken index: want error")
	}
}
