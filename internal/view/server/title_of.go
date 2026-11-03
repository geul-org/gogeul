//ff:func feature=server type=helper control=sequence
//ff:what titleOf — the page title of a tree path: the path with the Issuer ID and the document or codec ID shown by their names (the name address)
package server

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// titleOf is the title of the page at URL path p: the path without its
// leading /, with the Issuer ID replaced by its name in the Issuer list and
// `docs/<문서 ID>` or `codecs/<코덱 ID>` by the name the Issuer's document list
// gives it (the name address, namePath backwards), where they have one.
func titleOf(t resolve.Store, p string) string {
	segs := strings.Split(strings.TrimPrefix(p, wellKnown), "/")
	if len(segs) > 3 && (segs[1] == "docs" || segs[1] == "codecs") {
		if n, ok := documentLabels(t, segs[0])[segs[2]]; ok {
			segs = append([]string{segs[0], n}, segs[3:]...)
		}
	}
	if n, ok := issuerLabels(t)[segs[0]]; ok {
		segs[0] = n
	}
	return strings.TrimPrefix(wellKnown, "/") + strings.Join(segs, "/")
}
