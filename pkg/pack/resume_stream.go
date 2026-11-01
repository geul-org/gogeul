//ff:func feature=pack type=writer control=sequence
//ff:what ResumeStream — continue the stream in out from its last checkpoint: state, open objects found (part file or sealed object) and cut back to their checkpointed size, stray runs and temp files removed
package pack

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
	"slices"
)

// ResumeStream reopens the unfinished stream in out at its last Checkpoint
// (or the last point a stopped Close saved). The open data pack and index
// object are found (reopenPart: their part files, or the objects sealed from
// them after the checkpoint) and cut back to the bytes the checkpoint
// counted, runs spilled after it and temp files of a killed write are
// removed, and Count and Last tell the caller where to go on. o must carry the same codec settings as the stream
// was started with (the dictionary and Issuers come from the checkpoint), or
// the output differs from an uninterrupted run. A stream whose Close had
// begun acknowledges the caller's puts up to its last key without storing
// them; Close then goes on from where it stopped.
func ResumeStream(out string, o StreamOptions) (*Stream, error) {
	b, err := os.ReadFile(filepath.Join(out, streamDir, streamStateFile))
	if err != nil {
		return nil, err
	}
	s := &Stream{out: out, o: o.withDefaults()}
	if err := gob.NewDecoder(bytes.NewReader(b)).Decode(&s.st); err != nil {
		return nil, err
	}
	s.inum = issuerIndex(s.st.Issuers)
	if s.st.Dict.ID != 0 {
		s.dict = &Dict{ID: s.st.Dict.ID, Bytes: s.st.DictRaw}
	}
	if s.data, s.dsum, err = s.reopenPart(streamDataPart, kindData, s.st.Data); err != nil {
		return nil, err
	}
	if s.index, s.isum, err = s.reopenPart(streamIndexPart, kindIndex, s.st.Index); err != nil {
		return nil, err
	}
	if err := removeStrayRuns(out, slices.Concat(s.st.Runs, s.st.Close.Spent)); err != nil {
		return nil, err
	}
	if err := removeTemps(out); err != nil {
		return nil, err
	}
	return s, removeTemps(filepath.Join(out, "m"))
}
