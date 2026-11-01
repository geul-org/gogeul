//ff:func feature=pack type=writer control=sequence
//ff:what Stream.hWriteLeaf — write the pending H leaf block and append its head (first key, location) to the head file
package pack

// hWriteLeaf writes p.leaf as an H leaf block (addBlock) and its head to the
// head file, then empties the leaf.
func (s *Stream) hWriteLeaf(p *hPass) error {
	loc, err := s.addBlock(blockOf(blkHLeaf, p.leaf))
	if err != nil {
		return err
	}
	var h hEntry
	copy(h.key[:], p.leaf[0].Key)
	h.loc = loc
	var b [hEntrySize]byte
	p.w.Write(appendHEntry(b[:0], h)) // errors stay in w for its Flush
	p.count, p.since, p.leaf = p.count+1, p.since+1, p.leaf[:0]
	return nil
}
