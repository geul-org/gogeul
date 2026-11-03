//ff:func feature=server type=helper control=sequence
//ff:what issuerNames — Issuer name → ID from the tree's Issuer list (issuers/); empty when it is missing or broken
package server

import (
	"context"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// issuerNames reads the newest version of the Issuer list `issuers/` of the
// tree (latest.geul through its versions.geul; GEUL-ORG Issuer 이름).
func issuerNames(t resolve.Store) map[string]string {
	b, err := resolve.FetchAddr(context.Background(), resolve.Dir{Store: t}, resolve.Addr{Kind: resolve.AddrIssuers})
	if err != nil {
		return map[string]string{}
	}
	names, err := resolve.ParseIssuerNames(b)
	if err != nil {
		return map[string]string{}
	}
	return names
}
