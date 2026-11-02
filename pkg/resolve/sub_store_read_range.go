//ff:func feature=resolve type=resolver control=sequence
//ff:what subStore.ReadRange — pack.Store: a range of an object of the set
package resolve

import "context"

// ReadRange reads [off, off+n) of the object name of the set.
func (s subStore) ReadRange(ctx context.Context, name string, off, n uint64) ([]byte, error) {
	return s.base.ReadRange(ctx, s.object(name), off, n)
}
