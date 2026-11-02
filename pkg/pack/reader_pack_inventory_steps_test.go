//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderPackInventorySteps — a two-generation set yields replaced records twice, an h/-only record without a path, and P, named and h/ deletions with no location
package pack

import (
	"context"
	"errors"
	"slices"
	"testing"
)

func TestReaderPackInventorySteps(t *testing.T) {
	ctx := context.Background()
	out, _ := inventorySet(t, &Zstd{DictSize: -1})
	r := opened(t, out)
	var deleted []string
	var hashOnly, replaced []Loc
	shared := 0
	for _, e := range inventorySteps(t, r) {
		switch {
		case e.Deleted && e.Loc != (Loc{}):
			t.Fatalf("%+v", e)
		case e.Deleted:
			deleted = append(deleted, e.Path)
		case e.Path == "" && e.Hashed:
			hashOnly = append(hashOnly, e.Loc)
		case e.Path == pPath("0x00010000", 2, 1):
			replaced = append(replaced, e.Loc)
		}
		if e.Path != "" && e.Hashed {
			shared++
		}
	}
	slices.Sort(deleted)
	want := []string{pPath("0x00010000", 3, 1), docHPath(streamTestDoc(4, 1)), "n/b.geul"}
	slices.Sort(want)
	if !slices.Equal(deleted, want) || len(hashOnly) != 1 {
		t.Fatal(deleted, hashOnly)
	}
	if p, err := r.HashPath(ctx, hashOnly[0]); err != nil || p != docHPath(streamTestDoc(7, 1)) {
		t.Fatal(p, err)
	}
	if shared != 6 || len(replaced) != 2 || replaced[0].Obj >= replaced[1].Obj {
		t.Fatal(shared, replaced)
	}
	if _, err := r.Record(ctx, Loc{}); !errors.Is(err, ErrCorrupt) { // a deletion's Loc reads nothing
		t.Fatal(err)
	}
}
