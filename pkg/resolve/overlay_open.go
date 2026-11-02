//ff:func feature=resolve type=resolver control=sequence
//ff:what overlay.Open — fs.FS: the stored file or directory, else the loose one
package resolve

import (
	"errors"
	"io/fs"
)

// Open opens name.
func (o overlay) Open(name string) (fs.File, error) {
	f, err := o.stored.Open(name)
	if errors.Is(err, fs.ErrNotExist) {
		return o.loose.Open(name)
	}
	return f, err
}
