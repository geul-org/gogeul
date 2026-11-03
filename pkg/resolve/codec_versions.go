//ff:func feature=resolve type=resolver control=sequence
//ff:what CodecVersions — a codec document's version list: `<Issuer>/codecs/<코덱 ID>/versions.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// CodecVersions returns the version list of codec of issuer from src.
func CodecVersions(ctx context.Context, src Source, issuer []core.Word, codec uint32) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrCodec, Issuer: normalWords(issuer), ID: []core.Word{core.Word(codec >> 16), core.Word(codec)}, Versions: true})
}
