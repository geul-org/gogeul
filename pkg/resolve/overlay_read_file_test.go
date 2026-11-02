//ff:func feature=resolve type=resolver control=sequence
//ff:what TestOverlayReadFile — stored bytes win, a missing path is the loose file, then missing; other errors stop
package resolve

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestOverlayReadFile(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	d.write("a", []byte("loose"))
	d.write("b", []byte("loose"))
	o := overlay{stored: mapStore{fstest.MapFS{"a": {Data: []byte("stored")}}}, loose: dirStore{root: d.file("")}}
	if b, err := o.ReadFile("a"); err != nil || string(b) != "stored" {
		t.Fatal(string(b), err)
	}
	if b, err := o.ReadFile("b"); err != nil || string(b) != "loose" {
		t.Fatal(string(b), err)
	}
	if _, err := o.ReadFile("c"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if _, err := (overlay{stored: mapStore{fstest.MapFS{"dir/a": {Data: []byte("x")}}}, loose: o.loose}).ReadFile("dir"); err == nil || errors.Is(err, fs.ErrNotExist) {
		t.Fatal("a directory:", err)
	}
}
