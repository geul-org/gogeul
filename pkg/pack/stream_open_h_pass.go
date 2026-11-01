//ff:func feature=pack type=writer control=sequence
//ff:what Stream.openHPass — the H pass at its saved position: runs reopened after the entries taken, head file cut back to the saved heads, pending leaf and counts restored
package pack

import (
	"bufio"
	"os"
	"path/filepath"
	"slices"
)

// openHPass reopens the H pass where the state says it was (streamClose):
// heads written after the last save are dropped from the head file, as the
// index bytes after it are by ResumeStream.
func (s *Stream) openHPass() (*hPass, error) {
	c := s.st.Close
	rs, err := openRuns(s.out, s.st.Runs, c.Taken)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(filepath.Join(s.out, streamDir, streamHeadsFile), os.O_RDWR, 0)
	if err == nil {
		err = f.Truncate(c.HHeads * hEntrySize)
	}
	if err == nil {
		_, err = f.Seek(c.HHeads*hEntrySize, 0)
	}
	if err != nil {
		closeRuns(rs)
		if f != nil {
			f.Close()
		}
		return nil, err
	}
	return &hPass{rs: rs, heads: f, w: bufio.NewWriterSize(f, runBufSize), leaf: slices.Clone(c.HLeaf), n: c.HKeys, count: c.HHeads}, nil
}
