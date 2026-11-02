//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderFind — the newest layer holding a key decides; a deletion marker hides older layers; absent everywhere is NotFound; errors stop the search
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderFind(t *testing.T) {
	r, ctx := newReader(memStore{}, Options{}, &manifest{}, ""), context.Background()
	leaf := func(k string, l Loc) layer {
		return layer{blk: block{kind: blkHLeaf, keys: []string{k}, locs: []Loc{l}}}
	}
	ls := []layer{leaf("b", Loc{Enc: encGone}), leaf("a", Loc{Off: 70}), leaf("b", Loc{Off: 64})}
	if l, err := r.find(ctx, ls, "a", "p"); err != nil || l.Off != 70 {
		t.Fatal(l, err)
	}
	if _, err := r.find(ctx, ls, "b", "p"); !errors.Is(err, ErrNotFound) {
		t.Fatal("hidden", err)
	}
	if _, err := r.find(ctx, ls, "c", "p"); !errors.Is(err, ErrNotFound) {
		t.Fatal("absent", err)
	}
	bad := []layer{{blk: block{kind: blkHInner, keys: []string{"a"}, locs: []Loc{{Obj: 9}}}}}
	if _, err := r.find(ctx, bad, "a", "p"); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
