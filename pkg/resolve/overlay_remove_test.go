//ff:func feature=resolve type=resolver control=sequence
//ff:what TestOverlayRemove — a stored path leaves the stored side, a loose one the folder; a refused removal stays refused
package resolve

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestOverlayRemove(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.write("b", []byte("x"))
	loose := dirStore{root: d.file("")}
	s := localSite(d.Root)
	if err := (overlay{stored: s, loose: loose}).Remove("b"); err == nil {
		t.Fatal("a read-only stored side")
	}
	o := overlay{stored: mapStore{fstest.MapFS{"a": {Data: []byte("x")}}}, loose: loose}
	if err := o.Remove("a"); err != nil {
		t.Fatal(err)
	}
	if err := o.Remove("b"); err != nil {
		t.Fatal(err)
	}
	if _, err := loose.ReadFile("b"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if err := o.Remove("c"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
