//ff:func feature=server type=helper control=sequence
//ff:what namePath — a name address (<Issuer 이름>/<문서 이름>/…) as its tree path (<Issuer>/docs/<문서 ID>/… or codecs/<코덱 ID>/…); false when nothing was a name
package server

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// namePath reads URL path p as a name address (migration.md §3.8: the name
// address `geul.org/publication/` stays): the first segment as an Issuer name
// of the Issuer list, the segment after it as a name of the Issuer's document
// list, and returns the tree path with the IDs in their place — the document
// under `docs/<문서 ID>`, a codec under `codecs/<코덱 ID>`. It is false when
// neither is a known name.
func namePath(t resolve.Store, p string) (string, bool) {
	segs := strings.Split(strings.TrimPrefix(p, wellKnown), "/")
	issuer, named := issuerNames(t)[segs[0]]
	if named {
		segs[0] = issuer
	}
	if len(segs) > 1 {
		if id, ok := documentNames(t, segs[0])[segs[1]]; ok {
			segs = append([]string{segs[0], idFolder(id), id}, segs[2:]...)
			named = true
		}
	}
	return wellKnown + strings.Join(segs, "/"), named
}
