//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what invPrefixMatches — whether inventory entries begin with a checkpoint's entries (streamInv), field for field
package pack

// invPrefixMatches compares entries[i] with inv[i] for every i of inv
// (entries holds at least as many).
func invPrefixMatches(entries []invEntry, inv []streamInv) bool {
	for i, e := range inv {
		if entries[i] != (invEntry{kind: e.Kind, enc: e.Enc, flags: e.Flags, ref: e.Ref, doc: e.Doc, slot: e.Slot, n: e.N, off: e.Off}) {
			return false
		}
	}
	return true
}
