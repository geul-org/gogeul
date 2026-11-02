//ff:func feature=resolve type=resolver control=sequence
//ff:what TestOverlayWriteFile — writes go to the stored side and leave the folder alone
package resolve

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestOverlayWriteFile(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	m := mapStore{fstest.MapFS{}}
	o := overlay{stored: m, loose: dirStore{root: d.file("")}}
	if err := o.WriteFile("a", []byte("x")); err != nil {
		t.Fatal(err)
	}
	if b, err := m.ReadFile("a"); err != nil || string(b) != "x" {
		t.Fatal(err)
	}
	if _, err := o.loose.ReadFile("a"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
