//ff:func feature=toulmin type=encoder control=selection
//ff:what EncodeControl — encode a framing control as its GEUL ISSUER-format record
package toulmin

import (
	"errors"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// EncodeControl encodes p as its framing record: Edge Id, then a SPAN for
// PROSE, a 1-word section code for SECTION, or the case REF for CASE-BEGIN /
// CASE-END. A PROSE without a Span and a section code above one word are errors.
func EncodeControl(p ControlPacket, w Wire) ([]core.Word, error) {
	f, err := ControlCodec(p.Code)
	if err != nil {
		return nil, err
	}
	body := PutID(nil, p.ID)
	switch p.Code {
	case ControlProse:
		if p.Span == nil {
			return nil, errors.New("geul/toulmin: PROSE control requires a Span")
		}
		body, err = PutSpan(body, *p.Span, w.Span)
	case ControlSection:
		if p.Arg > 0xFFFF {
			return nil, fmt.Errorf("geul/toulmin: section code %d exceeds one word", p.Arg)
		}
		body = append(body, core.Word(p.Arg))
	default:
		body = PutID(body, TID(p.Arg))
	}
	if err != nil {
		return nil, err
	}
	return EdgeRecord(CodecRef{Codec: f}, body)
}
