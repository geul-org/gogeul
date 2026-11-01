//ff:func feature=pack type=helper control=sequence
//ff:what sumMatches — whether a file's bytes after the header, up to the open object's size, hash to its saved state
package pack

import (
	"bytes"
	"io"
)

// sumMatches compares the SHA-256 of f[headerSize:open.Off] with the digest
// of open.Sum.
func sumMatches(f io.ReaderAt, open streamOpen) (bool, error) {
	want, err := sumOfState(open.Sum)
	if err != nil {
		return false, err
	}
	h, err := partSum(f, open.Off)
	if err != nil {
		return false, err
	}
	return bytes.Equal(h.Sum(nil), want), nil
}
