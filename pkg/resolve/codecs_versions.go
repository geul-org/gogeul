//ff:func feature=resolve type=resolver control=sequence
//ff:what CodecsVersions — the version list of an Issuer's codec list: `<Issuer>/codecs/versions.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// CodecsVersions returns `<Issuer>/codecs/versions.geul` from src.
func CodecsVersions(ctx context.Context, src Source, issuer []core.Word) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrCodecs, Issuer: normalWords(issuer), Versions: true})
}
