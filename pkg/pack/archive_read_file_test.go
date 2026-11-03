//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveReadFile — staged bytes and removals win; sets otherwise; directories, absent and invalid names fail
package pack

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

func TestArchiveReadFile(t *testing.T) {
	a, _ := testArchive(t)
	if b, err := a.ReadFile("issuers/0x00000001.geul"); err != nil || len(b) == 0 {
		t.Fatal(err)
	}
	a.WriteFile("new.geul", []byte("n"))
	a.Remove("issuers/0x00000001.geul")
	if b, err := a.ReadFile("new.geul"); err != nil || string(b) != "n" {
		t.Fatal(err)
	}
	if _, err := a.ReadFile("issuers/0x00000001.geul"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if _, err := a.ReadFile("0x00000000"); err == nil || errors.Is(err, fs.ErrNotExist) {
		t.Fatal("dir", err)
	}
	if _, err := a.ReadFile("none"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if _, err := a.ReadFile("/abs"); !errors.Is(err, fs.ErrInvalid) {
		t.Fatal(err)
	}
	os.Remove(objPath(t, a.sets[0].r, a.sets[0].dir, kindData))
	if _, err := a.ReadFile("0x00000000/docs/0x00000001.geul"); err == nil {
		t.Fatal("broken")
	}
}
