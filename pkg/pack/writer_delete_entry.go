//ff:func feature=pack type=writer control=selection
//ff:what writer.deleteEntry — record one deletion in the current inventory (kind 5, kind 6 or a deleted hash)
package pack

// deleteEntry adds the deletion of path to w.inv.
func (w *writer) deleteEntry(path string) {
	ref := parsePath(path)
	switch ref.space {
	case spaceP:
		w.inv.entries = append(w.inv.entries, invEntry{kind: invPDel, ref: w.inum[ref.issuer], doc: ref.doc, slot: ref.slot})
	case spaceH:
		w.inv.deleted = append(w.inv.deleted, ref.hash)
	default:
		if ref.space == spaceNamed {
			path = ref.path
		}
		w.inv.entries = append(w.inv.entries, invEntry{kind: invNamedDel, ref: uint32(len(w.inv.paths))})
		w.inv.paths = append(w.inv.paths, path)
	}
}
