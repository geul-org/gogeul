//ff:func feature=delta type=decoder control=selection
//ff:what edgeID — the LOCAL Edge Id of an Edge record and the words after it; false for anything else
package delta

import "github.com/geul-org/gogeul/pkg/core"

// edgeReferenceNode is the REFERENCE-NODE Edge codec (edge/REFERENCE-NODE).
const edgeReferenceNode uint16 = 0x00

// edgeTriple is the TRIPLE Edge codec (edge/TRIPLE).
const edgeTriple uint16 = core.EdgeTriple

// edgeID splits an Edge record into its Edge Id (the leading LOCAL
// Reference, OVERVIEW Edge) and the rest of its payload.
func edgeID(r core.Record) (uint64, []core.Word, bool) {
	if r.Kind != core.KindEdge {
		return 0, nil, false
	}
	lead, n, err := core.DecodeRecord(r.Payload)
	id, lerr := core.DecodeLocalRef(lead)
	switch {
	case err != nil, lerr != nil, id == 0:
		return 0, nil, false
	default:
		return id, r.Payload[n:], true
	}
}
