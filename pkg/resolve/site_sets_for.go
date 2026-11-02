//ff:func feature=resolve type=resolver control=sequence
//ff:what site.setsFor — the sets that may serve a path: its Issuer's set then the top set for `<Issuer>/…`, every findable set otherwise
package resolve

import (
	"context"
	"strings"
)

// setsFor follows pack.Archive's read order: `<Issuer>/…` lives in that
// Issuer's set or the top set; anything else (h/, issuers.geul) may be in any.
func (s *site) setsFor(ctx context.Context, path string) ([]string, error) {
	first, _, nested := strings.Cut(path, "/")
	if nested && isIssuerID(first) {
		return []string{first, ""}, nil
	}
	return s.allSets(ctx, path)
}
