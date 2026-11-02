//ff:func feature=resolve type=resolver control=sequence
//ff:what site.WriteFile — Store: refused, a source on the GEULPROXY list is only read
package resolve

import "fmt"

// WriteFile refuses: documents are written to a data folder, not to a source.
func (s *site) WriteFile(name string, _ []byte) error {
	return fmt.Errorf("geul/resolve: %s: a GEULPROXY source is read-only", name)
}
