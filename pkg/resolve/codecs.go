//ff:func feature=resolve type=resolver control=sequence
//ff:what Codecs — an Issuer's codec list at version v (Latest for the newest): `<Issuer>/codecs/<v|latest>.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// Codecs returns issuer's codec list `<Issuer>/codecs/<버전 ID>.geul` from
// src, v Latest for `latest.geul`. issuer may be zero-padded (IssuerNormal).
func Codecs(ctx context.Context, src Source, issuer []core.Word, v uint32) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrCodecs, Issuer: normalWords(issuer), Version: v})
}
