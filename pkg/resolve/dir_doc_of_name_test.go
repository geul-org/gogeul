//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirDocOfName — a document or codec ID as is, a name by the document list; an unknown name and a broken index fail
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirDocOfName(t *testing.T) {
	d := issuersFixture(t)
	issuer := []core.Word{0, 0}
	if got, err := d.docOfName(issuer, "seven"); err != nil || core.FormatID(got) != "0xffff000000000007" {
		t.Fatalf("name: %v", err)
	}
	if got, err := d.docOfName(issuer, "0x0000000000000009"); err != nil || got[3] != 9 {
		t.Fatalf("id: %v", err)
	}
	if got, err := d.docOfName(issuer, "0x00000980"); err != nil || len(got) != 2 || got[1] != 0x0980 {
		t.Fatalf("codec id: %v %v", got, err)
	}
	if _, err := d.docOfName(issuer, "eight"); err == nil {
		t.Fatal("unknown: want error")
	}
	d.write(listPath(namesKey(geulIssuer)), []byte{1})
	if _, err := d.docOfName(issuer, "seven"); err == nil {
		t.Fatal("broken index: want error")
	}
}
