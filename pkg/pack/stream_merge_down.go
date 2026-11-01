//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.mergeDown — merge the H runs MergeRuns at a time into new runs until at most MergeRuns remain, saving after each merge
package pack

import "slices"

// mergeDown bounds the final merge: while there are more than MergeRuns
// runs, the first MergeRuns are merged into a new run (mergeRuns) put at the
// end of the list, the inputs are kept as Spent (removed with the working
// folder), and the state is saved, so a stop loses at most one merge. Each
// entry is merged about once; the order the runs are merged in does not
// change the result (hEntryLess is a total order).
func (s *Stream) mergeDown() error {
	for len(s.st.Runs) > s.o.MergeRuns {
		in := s.st.Runs[:s.o.MergeRuns]
		name := runName(s.st.RunCount)
		if err := mergeRuns(s.out, in, name); err != nil {
			return err
		}
		s.st.Close.Spent = append(s.st.Close.Spent, in...)
		s.st.Runs = append(slices.Clone(s.st.Runs[s.o.MergeRuns:]), name)
		s.st.RunCount++
		if err := s.saveState(); err != nil {
			return err
		}
	}
	return nil
}
