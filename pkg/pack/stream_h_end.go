//ff:func feature=pack type=writer control=sequence
//ff:what Stream.hEnd — the H layer from the pass: none without keys, the pending leaf alone as root, or that leaf written and the tree built over the head file
package pack

// hEnd finishes the H layer of p.n keys: with no written leaf the pending
// one is the root; otherwise it is written too and hTree groups the heads.
func (s *Stream) hEnd(p *hPass) ([]layer, error) {
	if p.n == 0 {
		return nil, nil
	}
	if p.count == 0 {
		return s.layerOf(blockOf(blkHLeaf, p.leaf), p.n)
	}
	if err := s.hWriteLeaf(p); err != nil {
		return nil, err
	}
	if err := p.w.Flush(); err != nil {
		return nil, err
	}
	root, err := s.hTree(p.heads, p.count)
	if err != nil {
		return nil, err
	}
	return s.layerOf(root, p.n)
}
