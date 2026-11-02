//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.readRange — a range of a listed object; a missing object is ErrCorrupt (the manifest names it), never NotFound
package pack

import (
	"context"
	"errors"
)

// readRange reads [off, off+n) of object o.
func (r *Reader) readRange(ctx context.Context, o object, off, n uint64) ([]byte, error) {
	b, err := r.s.ReadRange(ctx, objectName(o), off, n)
	if errors.Is(err, ErrNotFound) {
		return nil, corrupt(objectName(o) + " is listed in OBJECTS but missing")
	}
	return b, err
}
