//ff:func feature=server type=helper control=sequence
//ff:what documentNames — name → document ID of an Issuer's own documents (its document list); empty when missing
package server

import (
	"github.com/geul-org/gogeul/pkg/resolve"
)

// documentNames is docNameMaps' name → ID map; callers must not change it.
func documentNames(t resolve.Store, space string) map[string]string {
	return docNameMaps(t, space).byName
}
