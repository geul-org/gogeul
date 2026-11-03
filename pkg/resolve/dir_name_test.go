//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestDirName — names bind to published documents, again is a no-op, a second name appends, a taken name, a bad name, an unpublished document, a broken list or index fail
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDirName(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	issuer := []core.Word{0, 0}
	a, b := []core.Word{0xffff, 0, 0, 7}, []core.Word{0xffff, 0, 0, 8}
	if err := d.Name(issuer, a, "alpha"); err == nil {
		t.Fatal("unpublished: want error")
	}
	d.Publish(edgeDoc(7, "a"))
	d.Publish(edgeDoc(8, "b"))
	for _, n := range []string{"alpha", "alpha", "first"} {
		if err := d.Name(issuer, a, n); err != nil {
			t.Fatalf("%s: %v", n, err)
		}
	}
	got, _ := d.Documents(issuer)
	if len(got) != 2 || len(got[0].Names) != 2 || got[0].Names[1] != "first" {
		t.Fatalf("%+v", got)
	}
	if err := d.Name(issuer, b, "alpha"); err == nil {
		t.Fatal("taken: want error")
	}
	if err := d.Name(issuer, b, "Bad Name"); err == nil {
		t.Fatal("bad name: want error")
	}
	d.write(listPath(namesKey(geulIssuer)), []byte{1})
	if err := d.Name(issuer, b, "beta"); err == nil {
		t.Fatal("broken index: want error")
	}
	d.write(listPath(sampleKeyOf(8)), []byte{1})
	if err := d.Name(issuer, b, "beta"); err == nil {
		t.Fatal("broken list: want error")
	}
}
