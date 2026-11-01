//ff:func feature=pack type=writer control=sequence
//ff:what restorePart — a part file rebuilt from a sealed object: its bytes after the header up to off, the header left zero, synced
package pack

import (
	"io"
	"os"
)

// restorePart writes src[headerSize:off] to a fresh part file at dst (the
// header stays zero, as in a part file being written; closeObject fills it)
// and returns it open for writing. The sealed object is left in place.
func restorePart(dst, src string, off uint64) (*os.File, error) {
	in, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	defer in.Close()
	f, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(io.NewOffsetWriter(f, headerSize), io.NewSectionReader(in, headerSize, int64(off)-headerSize))
	if err == nil {
		err = f.Sync()
	}
	if err != nil {
		f.Close()
		return nil, err
	}
	return f, nil
}
