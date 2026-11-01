//ff:func feature=pack type=writer control=sequence
//ff:what Stream.hLayer — spill the held keys, merge the runs down to MergeRuns and start the H pass (first time only), then run the H pass from its saved position
package pack

// hLayer finishes the H layer (§4.7 kind 3 and 4). The first time it spills
// the held keys, merges the runs down to at most MergeRuns (mergeDown) and
// starts the pass (hStart); a resumed Close skips straight to the pass,
// which goes on from its last save (hPass). No key gives no layer.
func (s *Stream) hLayer() ([]layer, error) {
	if s.st.Close.Phase < closeHPass {
		if err := s.spill(); err != nil {
			return nil, err
		}
		if err := s.mergeDown(); err != nil {
			return nil, err
		}
		if err := s.hStart(); err != nil {
			return nil, err
		}
	}
	return s.hPass()
}
