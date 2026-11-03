//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirDeprecate — a deprecated version is skipped as newest but still served; twice, unknown and deleted versions fail
package resolve

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirDeprecate(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	issuer, doc := []core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}
	v1, v2 := edgeDoc(7, "a"), edgeDoc(7, "b")
	d.Publish(v1)
	d.Publish(v2)
	if err := d.Deprecate(issuer, doc, 2); err != nil {
		t.Fatal(err)
	}
	if b, _ := d.At(issuer, doc, 0); !bytes.Equal(b, v1) {
		t.Fatal("newest should skip the deprecated version")
	}
	if b, _ := d.At(issuer, doc, 2); !bytes.Equal(b, v2) {
		t.Fatal("deprecated version still served by number")
	}
	if err := d.Deprecate(issuer, doc, 2); err == nil {
		t.Fatal("twice: want error")
	}
	if err := d.Deprecate(issuer, doc, 5); err == nil {
		t.Fatal("unknown: want error")
	}
	d.Tombstone(issuer, doc, 1, 1)
	if err := d.Deprecate(issuer, doc, 1); err == nil {
		t.Fatal("deleted: want error")
	}
	d.write(listPath(sampleKeyOf(7)), []byte{1})
	if err := d.Deprecate(issuer, doc, 1); err == nil {
		t.Fatal("broken list: want error")
	}
}
