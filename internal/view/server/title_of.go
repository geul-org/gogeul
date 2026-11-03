//ff:func feature=server type=helper control=sequence
//ff:what titleOf — the page title of a tree path: the path with the Issuer ID shown by its name in the Issuer list
package server

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// titleOf is the title of the page at URL path p: the path without its
// leading /, with the Issuer ID replaced by its name in the Issuer list where
// it has one (namePath backwards). Document and codec IDs stay: they have no
// names (§1.1 101).
func titleOf(t resolve.Store, p string) string {
	segs := strings.Split(strings.TrimPrefix(p, wellKnown), "/")
	if n, ok := issuerLabels(t)[segs[0]]; ok {
		segs[0] = n
	}
	return strings.TrimPrefix(wellKnown, "/") + strings.Join(segs, "/")
}
