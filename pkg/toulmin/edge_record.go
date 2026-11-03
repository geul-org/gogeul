//ff:func feature=toulmin type=encoder control=sequence
//ff:what EdgeRecord — serialize a TANGEUL payload as an Edge record under its direct codec
package toulmin

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// EdgeRecord wraps payload as an Edge record under ref's direct TOULMIN codec:
// the TAG, then the payload.
func EdgeRecord(ref CodecRef, payload []core.Word) ([]core.Word, error) {
	return core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: ref.Codec, Payload: payload})
}
