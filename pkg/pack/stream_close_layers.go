//ff:func feature=pack type=writer control=sequence
//ff:what Stream.closeLayers — the steps of Close not done yet, each saved when done: last data pack, P layer, H layer, last index object
package pack

// closeLayers does what Close's saved progress lacks, in order, saving
// after each: the open data pack is closed (finishData), the P layer built
// (pLayer, kept in streamClose.P), the H layer built (hLayer, which saves
// within, kept in streamClose.H) and the open index object closed.
func (s *Stream) closeLayers() error {
	if s.st.Data.Num >= 0 {
		if err := s.finishData(); err != nil {
			return err
		}
		if err := s.closeSave(); err != nil {
			return err
		}
	}
	if s.st.Close.Phase < closePDone {
		pl, err := s.pLayer()
		if err != nil {
			return err
		}
		s.st.Close.P, s.st.Close.Phase = streamLayersOf(pl), closePDone
		if err := s.closeSave(); err != nil {
			return err
		}
	}
	if s.st.Close.Phase < closeHDone {
		hl, err := s.hLayer()
		if err != nil {
			return err
		}
		s.st.Close.H, s.st.Close.Phase = streamLayersOf(hl), closeHDone
		if err := s.closeSave(); err != nil {
			return err
		}
	}
	if s.st.Index.Num < 0 {
		return nil
	}
	if err := s.closeObject(s.index, s.st.Index, kindIndex, nil); err != nil {
		return err
	}
	s.index, s.st.Index, s.isum = nil, streamOpen{Num: -1}, nil
	return s.saveState()
}
