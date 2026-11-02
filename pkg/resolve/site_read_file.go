//ff:func feature=resolve type=resolver control=sequence
//ff:what site.ReadFile — fs.ReadFileFS: Fetch without a context, a missing path as fs.ErrNotExist
package resolve

import (
	"context"
	"errors"
	"io/fs"
)

// ReadFile reads name.
func (s *site) ReadFile(name string) ([]byte, error) {
	b, err := s.Fetch(context.Background(), name)
	if errors.Is(err, ErrNotFound) {
		return nil, &fs.PathError{Op: "read", Path: name, Err: fs.ErrNotExist}
	}
	return b, err
}
