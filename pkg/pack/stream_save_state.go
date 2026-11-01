//ff:func feature=pack type=writer control=sequence
//ff:what Stream.saveState — record the open objects' hash states, write the checkpoint (gob) to a temp file, sync it and rename it over the state file
package pack

import (
	"bytes"
	"encoding/gob"
	"path/filepath"
)

// saveState replaces out/.stream/state with the current streamState in one
// rename (writeFileSync), so a crash leaves either the old or the new
// checkpoint. The open objects' running hashes are saved with it
// (streamOpen.Sum); the caller has synced the bytes they cover.
func (s *Stream) saveState() error {
	s.st.Data.Sum, s.st.Index.Sum = hashState(s.dsum), hashState(s.isum)
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(s.st) // plain exported fields: never fails
	return writeFileSync(filepath.Join(s.out, streamDir, streamStateFile), b.Bytes())
}
