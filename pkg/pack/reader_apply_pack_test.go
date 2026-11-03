//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderApplyPack — entries then deleted hashes; an unreadable inventory or bad entry fails
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderApplyPack(t *testing.T) {
	ctx := context.Background()
	inv := inventory{issuers: []string{"0x00000000"}, entries: []invEntry{{kind: invPDel, doc: 8}}, deleted: []string{string(make([]byte, 32))}}
	obj, o := objBytes(kindData, nil, zstdFrame(encodeInventory(inv)))
	badInv := inv
	badInv.entries = []invEntry{{kind: 9}}
	obj2, o2 := objBytes(kindData, []byte("x"), zstdFrame(encodeInventory(badInv)))
	s := memStore{objectName(o): obj, objectName(o2): obj2}
	r := newReader(s, Options{}, &manifest{objects: []object{o, o2, {kind: kindData, size: 64}}}, "")
	st := map[string]Loc{"0x00000000/docs/0x0000000000000008/versions.geul": {}, hPath("0x00000000", string(make([]byte, 32))): {}, "x": {}}
	if err := r.applyPack(ctx, st, 0); err != nil || len(st) != 1 {
		t.Fatal(st, err)
	}
	if err := r.applyPack(ctx, st, 1); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	if err := r.applyPack(ctx, st, 2); err == nil {
		t.Fatal("missing object")
	}
}
