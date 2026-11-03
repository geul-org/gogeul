//ff:func feature=resolve type=resolver control=sequence
//ff:what DocVersions — a document's version list: `<Issuer>/docs/<문서 ID>/versions.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// DocVersions returns the version list of document doc of issuer from src.
func DocVersions(ctx context.Context, src Source, issuer []core.Word, doc uint64) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrDoc, Issuer: normalWords(issuer), ID: uintWords4(doc), Versions: true})
}
