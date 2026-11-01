//ff:func feature=pack type=writer control=sequence
//ff:what cutPart — hash a part file's bytes after the header up to off and cut it there; on error it is closed
package pack

import (
	"hash"
	"os"
)

// cutPart hashes f[headerSize:off] (partSum), so the stream goes on hashing
// the object, and truncates f to off. On error f is closed and nil returned.
func cutPart(f *os.File, off uint64) (*os.File, hash.Hash, error) {
	h, err := partSum(f, off)
	if err == nil {
		err = f.Truncate(int64(off))
	}
	if err != nil {
		f.Close()
		return nil, nil, err
	}
	return f, h, nil
}
