//ff:func feature=toulmin type=encoder control=sequence
//ff:what EncodeText — encode a TEXT record: Edge Id plus a UTF-8 Value
package toulmin

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// EncodeText encodes p as a TEXT record: its Edge Id, then the bytes as a
// UTF8-TEXT Value (even/odd codec, verbatim).
func EncodeText(p TextPacket, w Wire) ([]core.Word, error) {
	body, err := putValue(PutID(nil, p.TID), core.EncodeBytesValue(core.CodecUTF8TextEven, p.Bytes))
	if err != nil {
		return nil, err
	}
	return EdgeRecord(CodecRef{Codec: core.CodecEdgeIssuer2Word, Format: FormatText}, body)
}
