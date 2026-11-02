//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestOverlayOpen — the stored file, else the loose one, else missing; a refused open stays refused
package resolve

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestOverlayOpen(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.write("b", []byte("loose"))
	o := overlay{stored: mapStore{fstest.MapFS{"a": {Data: []byte("stored")}}}, loose: dirStore{root: d.file("")}}
	for _, n := range []string{"a", "b"} {
		f, err := o.Open(n)
		if err != nil {
			t.Fatal(n, err)
		}
		f.Close()
	}
	if _, err := o.Open("c"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	s, _ := served(t, t.TempDir())
	if _, err := (overlay{stored: s, loose: o.loose}).Open("b"); err == nil || errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
