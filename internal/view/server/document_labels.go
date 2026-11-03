//ff:func feature=server type=helper control=sequence
//ff:what documentLabels — document ID → name of an Issuer's own documents (its document list, the first name in name order)
package server

import "github.com/geul-org/gogeul/pkg/resolve"

// documentLabels is docNameMaps' ID → name map; callers must not change it.
func documentLabels(t resolve.Store, space string) map[string]string {
	return docNameMaps(t, space).byID
}
