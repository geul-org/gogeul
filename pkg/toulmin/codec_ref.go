//ff:type feature=toulmin type=model
//ff:what CodecRef — where a TANGEUL record type lives: a direct Edge codec or a GEUL ISSUER format number
package toulmin

import "github.com/geul-org/gogeul/pkg/core"

// CodecRef identifies a TANGEUL record type on the wire: either a direct TOULMIN
// Edge codec (Format == 0) or EDGE-ISSUER-2WORD (Codec 0x13D, TAG 0x1BDL) with a
// GEUL codec number (Format != 0).
type CodecRef struct {
	Codec  uint16
	Format uint32
}

// codecIssuer2Word is EDGE-ISSUER-2WORD, the codec of every GEUL codec number.
const codecIssuer2Word uint16 = core.CodecEdgeIssuer2Word
