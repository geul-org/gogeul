//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Archive.storedUnder — whether any set holds a live path under name that is not staged for removal; the caller holds a.mu
package pack

import "context"

// storedUnder asks each set (Reader.hasUnder); a set that fails counts as no.
func (a *Archive) storedUnder(ctx context.Context, name string) bool {
	for _, s := range a.sets {
		if s.r == nil {
			continue
		}
		if ok, err := s.r.hasUnder(ctx, name, a.dels); err == nil && ok {
			return true
		}
	}
	return false
}
