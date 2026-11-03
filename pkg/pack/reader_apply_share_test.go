//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderApplyShare — flags bit0 adds the h/ path; no flag adds nothing; an unreadable record fails
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderApplyShare(t *testing.T) {
	ctx := context.Background()
	doc := fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul")
	obj, o := objBytes(kindData, doc, nil)
	r := newReader(memStore{objectName(o): obj}, Options{}, &manifest{objects: []object{o}}, "")
	st := map[string]Loc{}
	l := Loc{Off: 64, Len: uint64(len(doc))}
	if err := r.applyShare(ctx, st, invEntry{}, l); err != nil || len(st) != 0 {
		t.Fatal(err)
	}
	if err := r.applyShare(ctx, st, invEntry{flags: flagHash}, l); err != nil || len(st) != 1 {
		t.Fatal(err)
	}
	if err := r.applyShare(ctx, st, invEntry{flags: flagHash}, Loc{Obj: 9}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
