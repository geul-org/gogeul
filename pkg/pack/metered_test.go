//ff:type feature=pack type=helper
//ff:what metered — test fixture: an Archive and the meter under it
package pack

// metered is an opened Archive and its store's read count.
type metered struct {
	a *Archive
	m *meterStore
}
