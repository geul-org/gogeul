//ff:func feature=pack type=writer control=sequence
//ff:what Stream.hTake — one merged H entry: a repeated key keeps its last record, a full leaf is written first, and the position is saved every saveLeaves leaves
package pack

// hTake adds e to the pass. Of entries with the same key the last
// (hEntryLess) wins, as it would when the inventories are applied in order.
// A leaf is written (hWriteLeaf) when the next key arrives after it filled,
// so a single leaf stays unwritten as the root. After saveLeaves written
// leaves the position is saved (hSave) with e already in the leaf.
func (s *Stream) hTake(p *hPass, e hEntry) error {
	k := string(e.key[:])
	if len(p.leaf) > 0 && p.leaf[len(p.leaf)-1].Key == k {
		p.leaf[len(p.leaf)-1].Loc = e.loc
		return nil
	}
	if len(p.leaf) >= s.o.Build.HLeaf {
		if err := s.hWriteLeaf(p); err != nil {
			return err
		}
	}
	p.leaf, p.n = append(p.leaf, streamKV{Key: k, Loc: e.loc}), p.n+1
	if p.since >= s.o.saveLeaves {
		return s.hSave(p)
	}
	return nil
}
