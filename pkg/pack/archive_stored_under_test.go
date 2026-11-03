//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveStoredUnder — a set without a generation is skipped; a set that holds the name answers; a set that fails counts as no
package pack

import (
	"context"
	"os"
	"testing"
)

func TestArchiveStoredUnder(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	a.sets = append([]*packSet{{}}, a.sets...)
	if !a.storedUnder(ctx, "0x00000000") || a.storedUnder(ctx, "zz") {
		t.Fatal("storedUnder")
	}
	s := a.sets[1]
	s.r.blocks = map[Loc]block{}
	os.Remove(objPath(t, s.r, s.dir, kindIndex))
	if a.storedUnder(ctx, "h") {
		t.Fatal("a broken set: want no")
	}
}
