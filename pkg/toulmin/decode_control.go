//ff:func feature=toulmin type=decoder control=sequence
//ff:what DecodeControl — decode a framing record (SECTION, CASE-BEGIN/END, PROSE)
package toulmin

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// DecodeControl decodes r as a framing record: the format number gives the
// control, then the Edge Id and the control's argument, which must use the
// whole payload.
func DecodeControl(r core.Record, w Wire) (ControlPacket, error) {
	ref, err := RecordRef(r)
	if err != nil {
		return ControlPacket{}, err
	}
	code, ok := ControlFromFormat(ref.Format)
	if !ok || ref.Codec != core.CodecEdgeIssuer2Word {
		return ControlPacket{}, fmt.Errorf("geul/toulmin: record 0x%03X/0x%08X is not a framing control", ref.Codec, ref.Format)
	}
	ws, pos := r.Payload, 0
	p := ControlPacket{Code: code}
	if p.ID, err = TakeID(ws, &pos); err != nil {
		return ControlPacket{}, err
	}
	if err := decodeControlArg(ws, &pos, &p, w); err != nil {
		return ControlPacket{}, err
	}
	return p, ExpectEnd(ws, pos, "control")
}
