//ff:func feature=resolve type=resolver control=sequence
//ff:what Hash — a document of an Issuer by content hash: `<Issuer>/hashes/<내용 해시>.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// Hash returns the document of issuer whose content hash is h from src. There
// is no hash address without an Issuer.
func Hash(ctx context.Context, src Source, issuer []core.Word, h [32]byte) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrHash, Issuer: normalWords(issuer), Hash: h})
}
