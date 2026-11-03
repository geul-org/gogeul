//ff:func feature=resolve type=resolver control=sequence
//ff:what site.allSets — the sets to search for a path without an Issuer: top, 0x00000000, those already open, then the place's own Issuers
package resolve

import (
	"context"
	"errors"
	"slices"
	"strings"
)

// allSets lists the top set, the GEUL Issuer's (which holds paths outside
// every Issuer, pack.Split), the sets opened so far and the Issuers the place
// shows. Only that last step costs anything: a folder lists them; a server
// cannot, so the IDs the Issuer list binds stand in, and a set of an Issuer it
// does not name is found once a `<Issuer>/…` read has opened it.
func (s *site) allSets(ctx context.Context, path string) ([]string, error) {
	more := s.opened()
	switch {
	case s.dir != "":
		ids, err := issuerFolders(s.dir)
		if err != nil {
			return nil, err
		}
		more = append(more, ids...)
	case !strings.HasPrefix(path, issuersDir+"/"):
		b, err := FetchAddr(ctx, s, Addr{Kind: AddrIssuers})
		if err != nil && !errors.Is(err, ErrNotFound) {
			return nil, err
		}
		more = append(more, issuerIDs(b)...)
	}
	slices.Sort(more)
	more = slices.DeleteFunc(slices.Compact(more), func(n string) bool { return n == "" || n == geulIssuer })
	return append([]string{"", geulIssuer}, more...), nil
}
