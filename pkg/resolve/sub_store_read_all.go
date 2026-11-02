//ff:func feature=resolve type=resolver control=sequence
//ff:what subStore.ReadAll — pack.Store: a whole object of the set
package resolve

import "context"

// ReadAll reads the object name of the set.
func (s subStore) ReadAll(ctx context.Context, name string) ([]byte, error) {
	return s.base.ReadAll(ctx, s.object(name))
}
