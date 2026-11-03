//ff:func feature=toulmin type=decoder control=sequence
//ff:what DecodeText — decode a TEXT record into its Edge Id and bytes
package toulmin

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// DecodeText decodes r as a TEXT record: the Edge Id then one UTF8-TEXT Value
// that ends the payload.
func DecodeText(r core.Record, w Wire) (TextPacket, error) {
	ref, err := RecordRef(r)
	if err != nil {
		return TextPacket{}, err
	}
	if ref.Codec != CodecText {
		return TextPacket{}, errors.New("geul/toulmin: record is not a TEXT")
	}
	ws, pos := r.Payload, 0
	id, err := TakeID(ws, &pos)
	if err != nil {
		return TextPacket{}, err
	}
	v, err := TakeValue(ws, &pos)
	if err != nil {
		return TextPacket{}, err
	}
	b, err := core.DecodeBytesValue(v, core.CodecUTF8TextEven)
	if err != nil {
		return TextPacket{}, err
	}
	return TextPacket{TID: id, Bytes: b}, ExpectEnd(ws, pos, "TEXT")
}
