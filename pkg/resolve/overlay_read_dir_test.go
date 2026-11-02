//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestOverlayReadDir — the union by name without a set's own files; either side alone is enough; both missing is missing; a real error stops
package resolve

import (
	"errors"
	"io/fs"
	"path/filepath"
	"slices"
	"testing"
	"testing/fstest"
)

func TestOverlayReadDir(t *testing.T) {
	d := Dir{Root: t.TempDir()}
	for _, n := range []string{"latest.txt", "m/x", "o/x", ".tmp-1", "0x0001/a", "0x0002/b", "only/c"} {
		d.write(n, []byte("x"))
	}
	stored := mapStore{fstest.MapFS{"0x0001/z": {Data: []byte("y")}, "0x0000/z": {Data: []byte("y")}, "here/z": {Data: []byte("y")}}}
	o := overlay{stored: stored, loose: dirStore{root: d.file("")}}
	ents, err := o.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, e := range ents {
		names = append(names, e.Name())
	}
	if !slices.Equal(names, []string{"0x0000", "0x0001", "0x0002", "here", "only"}) {
		t.Fatal(names)
	}
	for n, want := range map[string]string{"only": "c", "here": "z"} {
		if ents, err := o.ReadDir(n); err != nil || len(ents) != 1 || ents[0].Name() != want {
			t.Fatal(n, ents, err)
		}
	}
	if _, err := o.ReadDir("absent"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	s, _ := served(t, t.TempDir())
	if _, err := (overlay{stored: s, loose: o.loose}).ReadDir("."); err == nil {
		t.Fatal("a stored side that cannot list")
	}
	if _, err := (overlay{stored: stored, loose: dirStore{root: filepath.Join(d.file(""), "latest.txt")}}).ReadDir("."); err == nil || errors.Is(err, fs.ErrNotExist) {
		t.Fatal("a loose side that cannot list:", err)
	}
}
