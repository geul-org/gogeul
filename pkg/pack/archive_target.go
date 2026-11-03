//ff:func feature=pack type=writer control=sequence
//ff:what Archive.target — the set a staged path is written to: its holder; else its Issuer's set (top set if no Issuer sets layout); else the one Issuer set of this commit (h/ copies); else home
package pack

import (
	"context"
	"strings"
)

// target routes name. only is the single Issuer set receiving `<Issuer>/…`
// paths in this commit, or nil.
func (a *Archive) target(ctx context.Context, name string, only *packSet) (*packSet, error) {
	s, err := a.holder(ctx, name)
	if err != nil || s != nil {
		return s, err
	}
	first, _, nested := strings.Cut(name, "/")
	if isIssuer(first) && nested {
		if s := a.find(first); s != nil {
			return s, nil
		}
		if top := a.find(""); top != nil {
			return top, nil
		}
		return a.set(first), nil
	}
	return a.home(), nil
}
