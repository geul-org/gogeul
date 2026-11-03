//ff:func feature=resolve type=resolver control=sequence
//ff:what Issuers — the Issuer list at version v (Latest for the newest): `issuers/<v|latest>.geul`
package resolve

import "context"

// Issuers returns the Issuer list `issuers/<버전 ID>.geul` from src, v Latest
// (0) for `issuers/latest.geul`. The bytes are pure GEUL.
func Issuers(ctx context.Context, src Source, v uint32) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrIssuers, Version: v})
}
