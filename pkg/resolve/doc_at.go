//ff:func feature=resolve type=resolver control=sequence
//ff:what Doc — version v of a document (Latest for the newest): `<Issuer>/docs/<문서 ID>/<v|latest>.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// Doc returns version v of document doc (4-word ID) of issuer from src, v
// Latest for the newest live version.
func Doc(ctx context.Context, src Source, issuer []core.Word, doc uint64, v uint32) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrDoc, Issuer: normalWords(issuer), ID: uintWords4(doc), Version: v})
}
