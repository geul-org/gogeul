//ff:func feature=pack type=helper control=sequence
//ff:what meterStore.ReadRange — test fixture: count one range read and pass it on
package pack

import "context"

// ReadRange counts the read.
func (m *meterStore) ReadRange(ctx context.Context, name string, off, n uint64) ([]byte, error) {
	m.reads++
	return m.Store.ReadRange(ctx, name, off, n)
}
