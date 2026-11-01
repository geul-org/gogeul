//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.spill — sort the held H keys (key, then location) and write them as one run file
package pack

import (
	"path/filepath"
	"sort"
)

// spill writes the held H entries sorted by key then location (hEntryLess)
// to the next run file in the working folder (writeFileSync) and empties
// the buffer. Run
// boundaries do not change the output: the runs are merged when the stream
// closes.
func (s *Stream) spill() error {
	if len(s.hbuf) == 0 {
		return nil
	}
	sort.Slice(s.hbuf, func(i, j int) bool { return hEntryLess(s.hbuf[i], s.hbuf[j]) })
	name := runName(s.st.RunCount)
	buf := make([]byte, 0, len(s.hbuf)*hEntrySize)
	for _, e := range s.hbuf {
		buf = appendHEntry(buf, e)
	}
	if err := writeFileSync(filepath.Join(s.out, streamDir, name), buf); err != nil {
		return err
	}
	s.st.Runs, s.st.RunCount, s.hbuf = append(s.st.Runs, name), s.st.RunCount+1, s.hbuf[:0]
	return nil
}
