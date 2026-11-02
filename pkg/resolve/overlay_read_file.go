//ff:func feature=resolve type=resolver control=sequence
//ff:what overlay.ReadFile — fs.ReadFileFS: the stored bytes, else the loose file
package resolve

import (
	"errors"
	"io/fs"
)

// ReadFile reads name.
func (o overlay) ReadFile(name string) ([]byte, error) {
	b, err := o.stored.ReadFile(name)
	if errors.Is(err, fs.ErrNotExist) {
		return o.loose.ReadFile(name)
	}
	return b, err
}
