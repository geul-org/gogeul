//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveOpen — directories list, files read and seek (fstest.TestFS over the whole archive); invalid and absent names fail
package pack

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestArchiveOpen(t *testing.T) {
	a, _ := testArchive(t)
	if err := fstest.TestFS(a, "issuers/0x00000001.geul", "0x00000000/docs/0x0000000000000845/versions.geul"); err != nil {
		t.Fatal(err)
	}
	if _, err := a.Open("/x"); !errors.Is(err, fs.ErrInvalid) {
		t.Fatal(err)
	}
	if _, err := a.Open("none"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
}
