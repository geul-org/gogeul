//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.findPart — the open object's part file, restored from the object sealed from it when that happened after the checkpoint; neither matching is ErrCorrupt
package pack

import (
	"fmt"
	"os"
	"path/filepath"
)

// findPart opens the part file holding the checkpoint's open object (kind,
// open). An object sealed after the checkpoint (a pack cut while putting, or
// Close) was renamed into o/ and its part file is gone or holds a later
// object, so the objects in o/ the checkpoint does not know are tried first
// and a match is copied back (restorePart); otherwise the part file must
// match. Matching is partMatches.
func (s *Stream) findPart(part string, kind byte, open streamOpen) (*os.File, error) {
	names, err := s.unknownObjects(kind)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(s.out, streamDir, part)
	for _, n := range names {
		src := filepath.Join(s.out, "o", n)
		ok, err := s.partMatches(src, kind, open, true)
		if err != nil {
			return nil, err
		}
		if ok {
			return restorePart(path, src, open.Off)
		}
	}
	ok, err := s.partMatches(path, kind, open, false)
	if err == nil && !ok {
		err = fmt.Errorf("%w: %s: neither it nor a sealed object holds the checkpoint's %d bytes of object %d", ErrCorrupt, part, open.Off, open.Num)
	}
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_RDWR, 0)
}
