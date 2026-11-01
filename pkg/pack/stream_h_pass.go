//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.hPass — merge the (at most MergeRuns) runs into H leaves from the saved position, then finish the layer
package pack

// hPass takes every remaining entry in order (hTake: repeated keys, full
// leaves, periodic saves) and ends the layer (hEnd).
func (s *Stream) hPass() ([]layer, error) {
	p, err := s.openHPass()
	if err != nil {
		return nil, err
	}
	defer p.close()
	for {
		e, ok, err := nextMerged(p.rs)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		if err := s.hTake(p, e); err != nil {
			return nil, err
		}
	}
	return s.hEnd(p)
}
