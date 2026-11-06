//ff:func feature=toulmin type=decoder control=sequence
//ff:what RecordRef — the TANGEUL CodecRef of a decoded Edge record (direct codes only)
package toulmin

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// RecordRef returns the CodecRef of r. TANGEUL records are direct Edge records;
// an EDGE-ISSUER record is not one (its former GEUL formats are direct
// codes).
func RecordRef(r core.Record) (CodecRef, error) {
	if r.Kind != core.KindEdge {
		return CodecRef{}, fmt.Errorf("geul/toulmin: expected an Edge record, got kind %d", r.Kind)
	}
	if core.IssuerWidth(r.Kind, r.Codec) != 0 {
		return CodecRef{}, fmt.Errorf("geul/toulmin: EDGE-ISSUER record 0x%03X is not a TANGEUL record", r.Codec)
	}
	return CodecRef{Codec: r.Codec}, nil
}
