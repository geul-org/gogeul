//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderHashPath — the h/ path of a record by content hash; unreadable records and non-documents fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderHashPath(t *testing.T) {
	ctx := context.Background()
	doc := fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul")
	obj, o := objBytes(kindData, append(doc, 1, 2), nil)
	r := newReader(memStore{objectName(o): obj}, Options{}, &manifest{objects: []object{o}}, "")
	h, _ := contentHash(doc)
	if p, err := r.HashPath(ctx, Loc{Off: 64, Len: uint64(len(doc))}); err != nil || p != hPath("0x00000000", string(h[:])) {
		t.Fatal(p, err)
	}
	if _, err := r.HashPath(ctx, Loc{Off: 64 + uint64(len(doc)), Len: 2}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if _, err := r.HashPath(ctx, Loc{Obj: 9}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
