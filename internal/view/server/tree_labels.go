//ff:func feature=server type=helper control=iteration dimension=1
//ff:what treeLabels — the names to show in a listing: Issuer names, plus document names inside an Issuer
package server

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// treeLabels maps IDs to names for the listing of the directory at URL path p:
// Issuer IDs to their names in the Issuer list and, inside an Issuer, its
// document IDs to their names in its document list.
func treeLabels(t resolve.Store, p string) map[string]string {
	labels := issuerLabels(t)
	segs := strings.Split(strings.Trim(strings.TrimPrefix(p, wellKnown), "/"), "/")
	for id, n := range documentLabels(t, segs[0]) {
		labels[id] = n
	}
	return labels
}
