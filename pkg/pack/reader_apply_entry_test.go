//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderApplyEntry — sets, h/-only records, deletions, dictionaries ignored; unknown kinds and bad references fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderApplyEntry(t *testing.T) {
	ctx := context.Background()
	doc := fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul")
	obj, o := objBytes(kindData, doc, nil)
	r := newReader(memStore{objectName(o): obj}, Options{}, &manifest{objects: []object{o}}, "")
	inv := inventory{issuers: []string{"0x00000000"}, paths: []string{"a.geul"}}
	n := uint32(len(doc))
	st := map[string]Loc{}
	h, _ := contentHash(doc)
	steps := []invEntry{{kind: invP, doc: 8, slot: 1, n: n, off: 64}, {kind: invNamed, n: n, off: 64}, {kind: invH, n: n, off: 64}, {kind: invDict}}
	for _, e := range steps {
		if err := r.applyEntry(ctx, st, 0, inv, e); err != nil {
			t.Fatal(e, err)
		}
	}
	if len(st) != 3 || st[hPath("0x00000000", string(h[:]))].Len != uint64(n) {
		t.Fatal(st)
	}
	for _, e := range []invEntry{{kind: invPDel, doc: 8, slot: 1}, {kind: invNamedDel}} {
		if err := r.applyEntry(ctx, st, 0, inv, e); err != nil {
			t.Fatal(err)
		}
	}
	if len(st) != 1 {
		t.Fatal(st)
	}
	for _, e := range []invEntry{{kind: 9}, {kind: invP, ref: 3}, {kind: invPDel, ref: 3}, {kind: invH, off: 1}} {
		if err := r.applyEntry(ctx, st, 0, inv, e); !errors.Is(err, ErrCorrupt) {
			t.Fatal(e, err)
		}
	}
}
