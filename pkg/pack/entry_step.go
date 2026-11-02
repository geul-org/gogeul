//ff:func feature=pack type=resolver control=selection
//ff:what entryStep — the InventoryEntry of an entry of data pack obj; a dictionary gives none, unknown kinds and bad references are ErrCorrupt
package pack

import "fmt"

// entryStep translates e. The bool is false when there is no step to take.
func entryStep(obj uint32, inv inventory, e invEntry) (InventoryEntry, bool, error) {
	l := Loc{Obj: obj, Off: e.off, Len: uint64(e.n), Enc: e.enc}
	switch e.kind {
	case invP, invNamed:
		p, err := invPath(inv, e)
		return InventoryEntry{Path: p, Loc: l, Hashed: e.flags&flagHash != 0}, err == nil, err
	case invH:
		return InventoryEntry{Loc: l, Hashed: true}, true, nil
	case invPDel, invNamedDel:
		p, err := invPath(inv, e)
		return InventoryEntry{Path: p, Deleted: true}, err == nil, err
	case invDict:
		return InventoryEntry{}, false, nil
	default:
		return InventoryEntry{}, false, corrupt(fmt.Sprintf("inventory entry kind %d", e.kind))
	}
}
