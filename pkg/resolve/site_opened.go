//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what site.opened — the names of the sets open so far
package resolve

// opened lists the sets that exist among those tried.
func (s *site) opened() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	var out []string
	for n, r := range s.sets {
		if r != nil {
			out = append(out, n)
		}
	}
	return out
}
