//ff:func feature=pack type=helper control=sequence
//ff:what partSum — a running SHA-256 over an open object's bytes after the header, up to off
package pack

import (
	"crypto/sha256"
	"hash"
	"io"
)

// partSum hashes f[headerSize:off] (off ≥ headerSize), the bytes a stream
// hashes as it writes an object (streamOpen.Sum), and returns the hash to go
// on with.
func partSum(f io.ReaderAt, off uint64) (hash.Hash, error) {
	h := sha256.New()
	if _, err := io.Copy(h, io.NewSectionReader(f, headerSize, int64(off)-headerSize)); err != nil {
		return nil, err
	}
	return h, nil
}
