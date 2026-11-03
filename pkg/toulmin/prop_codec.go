//ff:func feature=toulmin type=encoder control=selection
//ff:what PropCodec — the record codec of a relation (do + once becomes DO-ONCE)
package toulmin

import "fmt"

// PropCodec returns the CodecRef an edge with PropCode p is written under (GEUL
// t002 TOULMIN 코덱). A `do` edge with once set uses DO-ONCE, which itself
// carries the once (the F_ONCE flag is then not written).
func PropCodec(p PropCode, once bool) (CodecRef, error) {
	switch p {
	case PropAttacks:
		return CodecRef{Codec: CodecAttacks}, nil
	case PropSupersedes:
		return CodecRef{Codec: CodecSupersedes}, nil
	case PropRun:
		return CodecRef{Codec: CodecRun}, nil
	case PropDo:
		return doCodec(once), nil
	case PropUndo:
		return CodecRef{Codec: CodecUndo}, nil
	case PropChecks:
		return CodecRef{Codec: CodecChecking}, nil
	case PropProvenance:
		return CodecRef{}, fmt.Errorf("geul/toulmin: provenance has no record layout yet")
	default:
		return propMovedRef(p)
	}
}
