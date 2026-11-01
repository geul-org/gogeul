//ff:func feature=pack type=writer control=sequence
//ff:what Stream.hStart — begin the H pass: an empty head file, nothing taken from any run, no pending leaf, saved
package pack

import (
	"os"
	"path/filepath"
)

// hStart empties the head file and sets the H pass position to the start
// (Phase closeHPass, Taken zero for each run), then saves.
func (s *Stream) hStart() error {
	if err := os.WriteFile(filepath.Join(s.out, streamDir, streamHeadsFile), nil, 0o644); err != nil {
		return err
	}
	c := &s.st.Close
	c.Phase, c.Taken, c.HLeaf, c.HKeys, c.HHeads = closeHPass, make([]int64, len(s.st.Runs)), nil, 0, 0
	return s.closeSave()
}
