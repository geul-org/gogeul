//ff:func feature=toulmin type=encoder control=sequence
//ff:what EncodeSource — encode the SOURCE record: Edge Id, SPAN width word, UTF-8 Value
package toulmin

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// EncodeSource encodes p as the SOURCE record — the authored document every
// SPAN anchors into: its Edge Id, the SPAN width word (code in bits 15-14, the
// rest 0) declaring w.Span, then the bytes as a UTF8-TEXT Value.
func EncodeSource(p TextPacket, w Wire) ([]core.Word, error) {
	code, err := core.WidthCode(w.Span)
	if err != nil {
		return nil, err
	}
	body, err := putValue(append(PutID(nil, p.TID), core.Word(uint16(code)<<14)), core.EncodeBytesValue(core.CodecUTF8TextEven, p.Bytes))
	if err != nil {
		return nil, err
	}
	return EdgeRecord(CodecRef{Codec: core.CodecEdgeIssuer2Word, Format: FormatSource}, body)
}
