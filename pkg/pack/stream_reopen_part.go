//ff:func feature=pack type=writer control=sequence
//ff:what Stream.reopenPart — the checkpoint's open object of kind as a part file cut back to its checkpointed size, with its running hash (nil when none was open)
package pack

import (
	"hash"
	"os"
)

// reopenPart finds the bytes of the object open at the checkpoint (findPart)
// and cuts the part file back to open.Off (cutPart), dropping bytes written
// after the checkpoint.
func (s *Stream) reopenPart(part string, kind byte, open streamOpen) (*os.File, hash.Hash, error) {
	if open.Num < 0 {
		return nil, nil, nil
	}
	f, err := s.findPart(part, kind, open)
	if err != nil {
		return nil, nil, err
	}
	return cutPart(f, open.Off)
}
