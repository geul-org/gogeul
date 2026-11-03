//ff:func feature=toulmin type=decoder control=sequence
//ff:what DecodeNode — decode a NODE record: type from the codec, header fields, slots and SPAN
package toulmin

import (
	"errors"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// DecodeNode decodes r as a NODE record (GEUL t002 TOULMIN NODE). The codec
// gives the node type; for the rule codecs the header role must equal the role
// the codec fixes. The reserved header bits must be 0, and the fields must use
// the whole payload.
func DecodeNode(r core.Record, w Wire) (NodePacket, error) {
	ref, err := RecordRef(r)
	if err != nil {
		return NodePacket{}, err
	}
	t, role, fixed, ok := NodeFromCodec(ref)
	if !ok {
		return NodePacket{}, fmt.Errorf("geul/toulmin: record 0x%03X is not a NODE", ref.Codec)
	}
	ws, pos := r.Payload, 0
	p := NodePacket{Type: t}
	if p.TID, err = TakeID(ws, &pos); err != nil {
		return NodePacket{}, err
	}
	header, err := TakeWord(ws, &pos, "NODE header")
	if err != nil {
		return NodePacket{}, err
	}
	p.Role, p.Level, p.Flags = RoleCode(header>>14), LevelCode((header>>9)&3), NodeFlags(header&0x087F)
	if header&0x3000 != 0 || fixed && p.Role != role {
		return NodePacket{}, errors.New("geul/toulmin: NODE header reserved bits set or role disagrees with codec")
	}
	if err := decodeNodeSlots(ws, &pos, &p, ObjKind((header>>7)&3), w); err != nil {
		return NodePacket{}, err
	}
	if p.Span, err = TakeSpan(ws, &pos, w.Span); err != nil {
		return NodePacket{}, err
	}
	return p, ExpectEnd(ws, pos, "NODE")
}
