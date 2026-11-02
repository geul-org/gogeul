//ff:func feature=resolve type=resolver control=sequence
//ff:what overlay.Remove — Store: remove from the stored side, else the loose file
package resolve

import (
	"errors"
	"io/fs"
)

// Remove deletes name.
func (o overlay) Remove(name string) error {
	err := o.stored.Remove(name)
	if errors.Is(err, fs.ErrNotExist) {
		return o.loose.Remove(name)
	}
	return err
}
