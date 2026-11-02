//ff:func feature=resolve type=helper control=sequence
//ff:what requestMeter.take — the requests and bytes since the last take
package resolve

// take returns and clears the record.
func (m *requestMeter) take() ([]string, int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	seen, sent := m.seen, m.sent
	m.seen, m.sent = nil, 0
	return seen, sent
}
