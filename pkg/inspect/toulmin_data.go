//ff:func feature=inspect type=decoder control=selection
//ff:what toulminData — decode a TOULMIN record to its packet and re-encode it for the lossless check
package inspect

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/toulmin"
)

// toulminData decodes r, of TOULMIN type ref, and returns the packet with the
// words it re-encodes to. A SOURCE updates the SPAN width in w. PREDICATE-
// DECLARE (whose condition layout belongs to the tangl expression AST) and any
// unknown type is an error.
func toulminData(r core.Record, ref toulmin.CodecRef, w *toulmin.Wire) (any, []core.Word, error) {
	_, _, _, isNode := toulmin.NodeFromCodec(ref)
	_, _, isLink := toulmin.PropFromCodec(ref)
	_, isControl := toulmin.ControlFromCodec(ref.Codec)
	switch {
	case ref.Codec == toulmin.CodecSource:
		return sourceData(r, w)
	case ref.Codec == toulmin.CodecText:
		p, err := toulmin.DecodeText(r, *w)
		enc, _ := toulmin.EncodeText(p, *w)
		return p, enc, err
	case isControl:
		p, err := toulmin.DecodeControl(r, *w)
		if err != nil {
			return nil, nil, err
		}
		enc, err := toulmin.EncodeControl(p, *w)
		return p, enc, err
	case isNode:
		p, err := toulmin.DecodeNode(r, *w)
		enc, _ := toulmin.EncodeNode(p, *w)
		return p, enc, err
	case isLink:
		p, err := toulmin.DecodeEdge(r, *w)
		enc, _ := toulmin.EncodeEdge(p, *w)
		return p, enc, err
	default:
		return nil, nil, fmt.Errorf("geul/inspect: codec 0x%03X has no typed form", ref.Codec)
	}
}
