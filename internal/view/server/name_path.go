//ff:func feature=server type=helper control=sequence
//ff:what namePath — an Issuer name address (<Issuer 이름>/…) as its tree path (<Issuer>/…); false when the first segment is not a known Issuer name
package server

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// namePath reads URL path p with its first segment as an Issuer name of the
// Issuer list (GEUL-ORG Issuer 이름) and returns the tree path with the
// Issuer ID in its place. Documents and codecs have no names (§1.1 101), so
// the rest of the path is kept as it is. It is false when the first segment
// is not a known Issuer name.
func namePath(t resolve.Store, p string) (string, bool) {
	segs := strings.Split(strings.TrimPrefix(p, wellKnown), "/")
	issuer, named := issuerNames(t)[segs[0]]
	if named {
		segs[0] = issuer
	}
	return wellKnown + strings.Join(segs, "/"), named
}
