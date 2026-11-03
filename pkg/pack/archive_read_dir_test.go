//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveReadDir — the top and Issuer listings as fs.ReadDir sees them; missing and invalid names fail; listing errors pass up
package pack

import (
	"errors"
	"io/fs"
	"testing"
)

func TestArchiveReadDir(t *testing.T) {
	a, _ := testArchive(t)
	top, err := fs.ReadDir(a, ".")
	if err != nil || len(top) != 2 || top[0].Name() != "0x00000000" || !top[0].IsDir() || top[1].Name() != "issuers" || !top[1].IsDir() {
		t.Fatal(top, err)
	}
	if es, err := a.ReadDir("0x00000000"); err != nil || len(es) != 3 || es[0].Name() != "codecs" {
		t.Fatal(es, err)
	}
	if _, err := a.ReadDir("none"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if _, err := a.ReadDir("../x"); !errors.Is(err, fs.ErrInvalid) {
		t.Fatal(err)
	}
	empty := &Archive{}
	if es, err := empty.ReadDir("."); err != nil || len(es) != 0 {
		t.Fatal(es, err)
	}
	a.idx = nil
	a.sets[0].r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	if _, err := a.ReadDir("."); err == nil {
		t.Fatal("broken")
	}
}
