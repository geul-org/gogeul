//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Archive.stagedUnder — whether a staged write lies under prefix; the caller holds a.mu
package pack

import "strings"

// stagedUnder looks through the staged writes, which are few.
func (a *Archive) stagedUnder(prefix string) bool {
	for p := range a.puts {
		if strings.HasPrefix(p, prefix) {
			return true
		}
	}
	return false
}
