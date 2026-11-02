//ff:func feature=resolve type=helper control=sequence
//ff:what bothSites — test fixture: one data folder read as a file:// source and as an https source
package resolve

import "testing"

// bothSites opens root twice.
func bothSites(t *testing.T, root string) []*site {
	t.Helper()
	s, _ := served(t, root)
	return []*site{localSite(root), s}
}
