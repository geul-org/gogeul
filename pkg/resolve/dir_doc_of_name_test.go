//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirDocOfName — an ID as is, a name by documents.geul; an unknown name and a broken index fail
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
	if _, err := d.docOfName(issuer, "eight"); err == nil {
		t.Fatal("unknown: want error")
	}
	d.write(listPath(namesKey(geulIssuer)), []byte{1})
	if _, err := d.docOfName(issuer, "seven"); err == nil {
		t.Fatal("broken index: want error")
	}
}
