//ff:func feature=toulmin type=encoder control=sequence
//ff:what EdgeRecord — serialize a TANGEUL payload as an Edge record under a direct codec or GEUL format
package toulmin

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// EdgeRecord wraps payload as an Edge record under ref: a direct TOULMIN codec,
// or EDGE-ISSUER-2WORD with the GEUL Issuer and ref's codec number.
func EdgeRecord(ref CodecRef, payload []core.Word) ([]core.Word, error) {
	r := core.Record{Kind: core.KindEdge, Codec: ref.Codec, Payload: payload}
	if ref.Format != 0 {
		r.Issuer = []core.Word{core.Word(IssuerGEUL >> 16), core.Word(IssuerGEUL)}
		r.Format = ref.Format
	}
	return core.EncodeRecord(r)
}
