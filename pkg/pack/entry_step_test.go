//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestEntryStep — each entry kind becomes its step with path, location and flags; dictionaries give none; unknown kinds and bad references fail
package pack

import (
	"errors"
	"testing"
)

func TestEntryStep(t *testing.T) {
	inv := sampleInventory()
	p := pPath("0x00000000", 8, 1)
	cases := []struct {
		e    invEntry
		want InventoryEntry
	}{
		{invEntry{kind: invP, enc: encZstd, flags: flagHash, doc: 8, slot: 1, n: 3, off: 64}, InventoryEntry{Path: p, Loc: Loc{Obj: 5, Off: 64, Len: 3, Enc: encZstd}, Hashed: true}},
		{invEntry{kind: invP, doc: 8, slot: 1, n: 3, off: 64}, InventoryEntry{Path: p, Loc: Loc{Obj: 5, Off: 64, Len: 3}}},
		{invEntry{kind: invNamed, n: 2, off: 67}, InventoryEntry{Path: "issuers/0x00000001.geul", Loc: Loc{Obj: 5, Off: 67, Len: 2}}},
		{invEntry{kind: invH, enc: encZstd, n: 4, off: 70}, InventoryEntry{Loc: Loc{Obj: 5, Off: 70, Len: 4, Enc: encZstd}, Hashed: true}},
		{invEntry{kind: invPDel, doc: 8, slot: 1}, InventoryEntry{Path: p, Deleted: true}},
		{invEntry{kind: invNamedDel}, InventoryEntry{Path: "issuers/0x00000001.geul", Deleted: true}},
	}
	for _, c := range cases {
		if got, ok, err := entryStep(5, inv, c.e); err != nil || !ok || got != c.want {
			t.Fatalf("%+v: %+v %v %v", c.e, got, ok, err)
		}
	}
	if _, ok, err := entryStep(5, inv, invEntry{kind: invDict, ref: 40000, n: 9, off: 64}); ok || err != nil {
		t.Fatal(ok, err)
	}
	for _, e := range []invEntry{{kind: 9}, {kind: invP, ref: 3}, {kind: invNamed, ref: 3}, {kind: invPDel, ref: 3}, {kind: invNamedDel, ref: 3}} {
		if _, ok, err := entryStep(5, inv, e); ok || !errors.Is(err, ErrCorrupt) {
			t.Fatal(e, ok, err)
		}
	}
}
