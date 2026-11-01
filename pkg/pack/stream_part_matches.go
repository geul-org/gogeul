//ff:func feature=pack type=writer control=selection
//ff:what Stream.partMatches — whether a file holds the checkpoint's open object up to its size: the saved hash, or for an older checkpoint its P leaves (index) or inventory (sealed data pack)
package pack

import (
	"errors"
	"io/fs"
	"os"
)

// partMatches tells whether the file at path (a part file, or a sealed
// object when sealed) begins with the bytes of the open object (kind, open)
// up to open.Off. A missing or shorter file does not. With a saved hash
// state the bytes after the header must hash to it. A checkpoint from before
// the hash (no Sum) is checked by structure: an index object by the P leaf
// blocks the checkpoint wrote to it (legacyIndexMatches), a sealed data pack
// by its inventory (legacyDataMatches); a data part file is taken as it is.
func (s *Stream) partMatches(path string, kind byte, open streamOpen, sealed bool) (bool, error) {
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil || uint64(fi.Size()) < open.Off {
		return false, err
	}
	switch {
	case len(open.Sum) > 0:
		return sumMatches(f, open)
	case kind == kindIndex:
		return legacyIndexMatches(f, open, s.st.PHeads), nil
	case sealed:
		return legacyDataMatches(f, open, s.st.Inv, s.st.Paths), nil
	}
	return true, nil
}
