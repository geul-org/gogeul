//ff:func feature=server type=helper control=sequence
//ff:what treeLabels — the names to show in a listing: Issuer names at the top of the tree, none elsewhere
package server

import "github.com/geul-org/gogeul/pkg/resolve"

// treeLabels maps IDs to names for the listing of the directory at URL path p:
// at the top of the tree, Issuer IDs to their names in the Issuer list. Below
// it nothing is named: documents and codecs have no names (§1.1 101), and an
// 8-digit codec ID must not take an Issuer's name.
func treeLabels(t resolve.Store, p string) map[string]string {
	if treeName(p) != "." {
		return map[string]string{}
	}
	return issuerLabels(t)
}
