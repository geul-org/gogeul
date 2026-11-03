//ff:func feature=resolve type=resolver control=sequence
//ff:what DocsVersions — the version list of an Issuer's document list: `<Issuer>/docs/versions.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// DocsVersions returns `<Issuer>/docs/versions.geul` from src.
func DocsVersions(ctx context.Context, src Source, issuer []core.Word) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrDocs, Issuer: normalWords(issuer), Versions: true})
}
