//ff:func feature=resolve type=resolver control=sequence
//ff:what Docs — an Issuer's document list at version v (Latest for the newest): `<Issuer>/docs/<v|latest>.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// Docs returns issuer's document list `<Issuer>/docs/<버전 ID>.geul` from src,
// v Latest for `latest.geul`.
func Docs(ctx context.Context, src Source, issuer []core.Word, v uint32) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrDocs, Issuer: normalWords(issuer), Version: v})
}
