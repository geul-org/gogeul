//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirIssuerOfName — an ID as is, a name by issuers.geul; no index, a broken index and an unknown name fail
package resolve

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirIssuerOfName(t *testing.T) {
	d := issuersFixture(t)
	if got, err := d.issuerOfName("0x00000001"); err != nil || !slices.Equal(got, []core.Word{0, 1}) {
		t.Fatalf("id: %v", err)
	}
	if got, err := d.issuerOfName("geul.org"); err != nil || !slices.Equal(got, []core.Word{0, 0}) {
		t.Fatalf("name: %v", err)
	}
	if _, err := d.issuerOfName("nobody"); err == nil {
		t.Fatal("unknown: want error")
	}
	d.write(issuersPath(), []byte{1})
	if _, err := d.issuerOfName("geul.org"); err == nil {
		t.Fatal("broken index: want error")
	}
	if _, err := (Dir{Root: t.TempDir()}).issuerOfName("geul.org"); err == nil {
		t.Fatal("no index: want error")
	}
}
