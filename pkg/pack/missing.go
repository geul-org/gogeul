//ff:func feature=pack type=helper control=sequence
//ff:what missing — a file-not-found error as ErrNotFound, other errors unchanged
package pack

import (
	"errors"
	"io/fs"
)

// missing maps fs.ErrNotExist to notFound(name).
func missing(err error, name string) error {
	if errors.Is(err, fs.ErrNotExist) {
		return notFound(name)
	}
	return err
}
