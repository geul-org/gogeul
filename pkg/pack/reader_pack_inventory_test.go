//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderPackInventory — steps in order with deleted hashes last, dictionaries skipped; fn's error stops the walk; only a readable data pack and valid entries are handed out
package pack

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestReaderPackInventory(t *testing.T) {
	ctx := context.Background()
	inv := sampleInventory()
	inv.entries = append(inv.entries[:2:2], invEntry{kind: invDict, ref: 40000, n: 5, off: 70}, inv.entries[2])
	obj, o := objBytes(kindData, []byte("abcde"), zstdFrame(encodeInventory(inv)))
	idx, io := objBytes(kindIndex, nil, nil)
	gone, gno := objBytes(kindData, nil, zstdFrame(encodeInventory(inventory{deleted: inv.deleted})))
	bad, bo := objBytes(kindData, []byte("x"), zstdFrame(encodeInventory(inventory{entries: []invEntry{{kind: 9}}})))
	broken, ko := objBytes(kindData, nil, []byte{1, 2})
	s := memStore{objectName(o): obj, objectName(io): idx, objectName(gno): gone, objectName(bo): bad, objectName(ko): broken}
	r := newReader(s, Options{}, &manifest{objects: []object{io, o, gno, bo, ko}}, "")
	var got []InventoryEntry
	collect := func(e InventoryEntry) error { got = append(got, e); return nil }
	if err := r.PackInventory(ctx, 1, collect); err != nil {
		t.Fatal(err)
	}
	want := []InventoryEntry{
		{Path: pPath("0x0000", 8, 1), Loc: Loc{Obj: 1, Off: 64, Len: 3}, Hashed: true},
		{Path: "issuers.geul", Loc: Loc{Obj: 1, Off: 67, Len: 2}},
		{Path: pPath("0x0000", 9, 0), Deleted: true},
		{Path: hPath(strings.Repeat("d", 32)), Deleted: true},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("%+v", got)
	}
	stop, calls := errors.New("stop"), 0
	halt := func(InventoryEntry) error { calls++; return stop }
	if err := r.PackInventory(ctx, 1, halt); err != stop || calls != 1 {
		t.Fatal(calls, err)
	}
	if err := r.PackInventory(ctx, 2, halt); err != stop || calls != 2 {
		t.Fatal(calls, err)
	}
	got = nil
	for _, n := range []uint32{0, 5, 0xFFFFFFFF} { // an index object, past OBJECTS: the caller's mistake, not corruption
		if err := r.PackInventory(ctx, n, collect); err == nil || errors.Is(err, ErrCorrupt) {
			t.Fatal(n, err)
		}
	}
	for _, n := range []uint32{3, 4} { // a bad entry, a broken frame
		if err := r.PackInventory(ctx, n, collect); !errors.Is(err, ErrCorrupt) {
			t.Fatal(n, err)
		}
	}
	if len(got) != 0 {
		t.Fatalf("%+v", got)
	}
}
