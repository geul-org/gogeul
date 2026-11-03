//ff:func feature=resolve type=resolver control=sequence
//ff:what TestDirLog — versions in order with publication times and NEXT-LOCAL-ID; an unpublished document is ErrNotFound; a broken list errors
package resolve

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirLog(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	issuer, doc := []core.Word{0, 0}, []core.Word{0xffff, 0, 0, 7}
	if _, err := d.Log(issuer, doc); !errors.Is(err, ErrNotFound) {
		t.Fatalf("unpublished: %v", err)
	}
	d.Publish(edgeDoc(7, "a"))
	d.Publish(edgeDoc(7, "a", "b"))
	d.Deprecate(issuer, doc, 1)
	h, err := d.Log(issuer, doc)
	if err != nil || len(h.Entries) != 2 || h.Next != 3 || !h.Entries[0].Deprecated || h.Entries[1].Version != 2 {
		t.Fatalf("%+v %v", h, err)
	}
	if h.Entries[0].At.IsZero() || h.Entries[1].At.IsZero() {
		t.Fatalf("times: %+v", h.Entries)
	}
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	d.write(listPath(k), []byte{1})
	if _, err := d.Log(issuer, doc); err == nil {
		t.Fatal("broken list: want error")
	}
}
