//ff:type feature=pack type=model
//ff:what InventoryEntry — one step of replaying a data pack: a record or a deletion, its path, location and h/ sharing
package pack

// InventoryEntry is one step of a data pack's inventory (§4.5, §7.1). Steps
// apply in order and later ones win: a record replaces any earlier record of
// its path, a deletion removes it, also across packs.
//
// Path is the P or named path, or the `h/` path of a deleted hash; it is ""
// for an h/-only record. Loc is zero for a deletion. Hashed means the record
// also serves `h/<content hash>.geul`. Only the record's bytes name that path,
// so it is not read here: HashPath(Loc) gives it.
type InventoryEntry struct {
	Path    string
	Loc     Loc
	Deleted bool
	Hashed  bool
}
