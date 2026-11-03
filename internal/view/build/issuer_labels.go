//ff:func feature=build type=helper control=sequence
//ff:what issuerLabels — Issuer ID → bound name from the tree's Issuer list (issuers/); empty when it is missing or broken
package build

import "github.com/geul-org/gogeul/pkg/resolve"

// issuerLabels inverts the names of the Issuer list at the top of the tree.
func issuerLabels(t resolve.Store) map[string]string {
	return invertNames(issuerNames(t))
}
