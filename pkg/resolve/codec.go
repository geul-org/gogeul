//ff:func feature=resolve type=resolver control=sequence
//ff:what Codec — version v of a codec document (Latest for the newest): `<Issuer>/codecs/<코덱 ID>/<v|latest>.geul`
package resolve

import (
	"context"

	"github.com/geul-org/gogeul/pkg/core"
)

// Codec returns version v of codec (the 2-word codec ID; a GEUL direct code's is
// its TAG with L=0) of issuer from src, v Latest for the newest live version.
func Codec(ctx context.Context, src Source, issuer []core.Word, codec uint32, v uint32) ([]byte, error) {
	return FetchAddr(ctx, src, Addr{Kind: AddrCodec, Issuer: normalWords(issuer), ID: []core.Word{core.Word(codec >> 16), core.Word(codec)}, Version: v})
}
