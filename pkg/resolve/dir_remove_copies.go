//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.removeCopies — deletes the h/ copies of the given hashes, keeping some and ignoring missing ones
package resolve

import (
	"errors"
	"io/fs"
)

// removeCopies deletes `<issuer>/hashes/<hash>.geul` for each hash in hs that is not in
// keep. A copy already gone is not an error.
func (d Dir) removeCopies(issuer string, hs [][32]byte, keep map[[32]byte]bool) error {
	for _, h := range hs {
		if keep[h] {
			continue
		}
		if err := d.store().Remove(hashPath(issuer, h)); err != nil && !errors.Is(err, fs.ErrNotExist) {
			return err
		}
	}
	return nil
}
