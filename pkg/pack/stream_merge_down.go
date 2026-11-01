//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.mergeDown — merge the H runs mergeWidth at a time into new runs until at most mergeWidth remain, saving after each merge, a safe point before and within each
package pack

import "slices"

// mergeDown bounds the final merge: while there are more runs than the merge
// may read at once (mergeWidth, asked again for each merge), that many runs
// are merged into a new run (mergeRuns) put at the end of the list, the
// inputs are kept as Spent (removed with the working folder), and the state
// is saved, so a stop loses at most one merge: the unfinished run is a stray
// run to ResumeStream. Each entry is merged about once at the full width; the
// order and width of the merges do not change the result (hEntryLess is a
// total order).
func (s *Stream) mergeDown() error {
	pace := func() error { return s.pace("close:merge") }
	for w := s.mergeWidth(); len(s.st.Runs) > w; w = s.mergeWidth() {
		if err := pace(); err != nil {
			return err
		}
		in := s.st.Runs[:w]
		name := runName(s.st.RunCount)
		if err := mergeRuns(s.out, in, name, pace); err != nil {
			return err
		}
		s.st.Close.Spent = append(s.st.Close.Spent, in...)
		s.st.Runs = append(slices.Clone(s.st.Runs[w:]), name)
		s.st.RunCount++
		if err := s.saveState(); err != nil {
			return err
		}
	}
	return nil
}
