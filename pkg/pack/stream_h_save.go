//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.hSave — save the H pass position: head file flushed and synced, entries taken per run, pending leaf, counts, open objects synced with the state
package pack

import "slices"

// hSave makes the pass durable: the heads written so far are flushed and
// synced, the pass position goes into streamClose, and closeSave syncs the
// index object and saves the state.
func (s *Stream) hSave(p *hPass) error {
	if err := p.w.Flush(); err != nil {
		return err
	}
	if err := p.heads.Sync(); err != nil {
		return err
	}
	c := &s.st.Close
	c.Taken = c.Taken[:0]
	for _, rr := range p.rs {
		c.Taken = append(c.Taken, rr.taken)
	}
	c.HLeaf, c.HKeys, c.HHeads, p.since = slices.Clone(p.leaf), p.n, p.count, 0
	return s.closeSave()
}
