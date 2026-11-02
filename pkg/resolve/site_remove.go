//ff:func feature=resolve type=resolver control=sequence
//ff:what site.Remove — Store: refused, a source on the GEULPROXY list is only read
package resolve

import "fmt"

// Remove refuses: documents are removed from a data folder, not from a source.
func (s *site) Remove(name string) error {
	return fmt.Errorf("geul/resolve: %s: a GEULPROXY source is read-only", name)
}
