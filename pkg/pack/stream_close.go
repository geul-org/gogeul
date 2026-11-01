//ff:func feature=pack type=writer control=sequence
//ff:what Stream.Close — save the start, finish the last data pack, the P and H layers and the index objects (saving progress after each), write the manifest and latest.txt, remove the working folder
package pack

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Close ends the stream (§6 steps 4–6): the open data pack gets its
// inventory, the P layer (pLayer) and the H layer (hLayer) are finished, the
// open index object is closed, and the manifest of generation 1 (no previous
// manifest) is written, then latest.txt, last. The working folder is removed.
// A stream with no record is an error. Close saves its progress as it goes
// (closeBegin, closeLayers): when it stops (an error, or the process is
// killed) ResumeStream and Close again go on from the last save and write
// the same bytes; a stop between latest.txt and the removal only repeats
// the last writes. Its memory does not grow with the number of H runs.
func (s *Stream) Close() (BuildResult, error) {
	if s.st.Count == 0 && len(s.st.Named) == 0 {
		return BuildResult{}, errors.New("geul/pack: a stream with no record makes no set")
	}
	named, err := sortedNamed(s.st.Named)
	if err == nil {
		err = s.closeBegin()
	}
	if err == nil {
		err = s.closeLayers()
	}
	if err != nil {
		return BuildResult{}, err
	}
	mb := encodeManifest(s.manifest(named, layersOf(s.st.Close.P), layersOf(s.st.Close.H)))
	sum := sha256.Sum256(mb)
	name := hex.EncodeToString(sum[:])
	if err := writeFileSync(filepath.Join(s.out, "m", name+".gpm"), mb); err != nil {
		return BuildResult{}, err
	}
	if err := writeFileSync(filepath.Join(s.out, latestFile), []byte(fmt.Sprintf("%d %s\n", s.st.Gen, name))); err != nil {
		return BuildResult{}, err
	}
	res := s.st.Res
	res.Generation, res.Manifest, res.Changed, res.Bytes = s.st.Gen, name, true, res.Bytes+int64(len(mb))
	return res, os.RemoveAll(filepath.Join(s.out, streamDir))
}
