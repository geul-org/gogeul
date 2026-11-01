//ff:func feature=pack type=resolver control=sequence
//ff:what Archive.isDir — whether any live path lies under name (the top always is a directory), answered from staged writes and key ranges without listing
package pack

import "context"

// isDir reports a directory; read errors count as no. A file-shaped name has
// no key range, so "no such file" costs no index walk (ReadFile).
func (a *Archive) isDir(name string) bool {
	if name == "." {
		return true
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.stagedUnder(dirPrefix(name)) || a.storedUnder(context.Background(), name)
}
