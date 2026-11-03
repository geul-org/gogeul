//ff:func feature=resolve type=resolver control=sequence
//ff:what IssuersVersions — the version list of the Issuer list: `issuers/versions.geul`
package resolve

import "context"

// IssuersVersions returns `issuers/versions.geul` from src.
func IssuersVersions(ctx context.Context, src Source) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrIssuers, Versions: true})
}
