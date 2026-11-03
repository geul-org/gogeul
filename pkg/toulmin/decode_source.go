//ff:func feature=toulmin type=decoder control=sequence
//ff:what DecodeSource — decode the SOURCE record into its text and the declared SPAN width
package toulmin

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/core"
)

// DecodeSource decodes r as the SOURCE record and returns its text plus the
// SPAN width it declares for every other record. The width word's low 14 bits are reserved and must be 0.
func DecodeSource(r core.Record) (TextPacket, int, error) {
	ref, err := RecordRef(r)
	if err != nil {
		return TextPacket{}, 0, err
	}
	if ref.Codec != CodecSource {
		return TextPacket{}, 0, errors.New("geul/toulmin: record is not the SOURCE")
	}
	ws, pos := r.Payload, 0
	id, err := TakeID(ws, &pos)
	if err != nil {
		return TextPacket{}, 0, err
	}
	sw, err := TakeWord(ws, &pos, "SOURCE SPAN width")
	if err != nil || sw&0x3FFF != 0 {
		return TextPacket{}, 0, errors.New("geul/toulmin: SOURCE SPAN width word missing or reserved bits set")
	}
	v, err := TakeValue(ws, &pos)
	if err != nil {
		return TextPacket{}, 0, err
	}
	b, err := core.DecodeBytesValue(v, core.CodecUTF8TextEven)
	if err != nil {
		return TextPacket{}, 0, err
	}
	return TextPacket{TID: id, Bytes: b}, core.WidthWords(uint8(sw >> 14)), ExpectEnd(ws, pos, "SOURCE")
}
