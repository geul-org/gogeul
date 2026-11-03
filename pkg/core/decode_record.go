//ff:func feature=core type=decoder control=sequence
//ff:what DecodeRecord — decode one GEUL record from the head of ws, returning it and the words consumed
package core

import (
	"errors"
	"fmt"
)

// DecodeRecord decodes the record at the head of ws. It enforces the 000 prefix
// and the canonical Length, reads the Kind from the section (KindOf), slices the Issuer
// and the 2-word codec number off an ISSUER codec, and returns the Record with
// the total word count, so a caller can skip records without a codebook.
func DecodeRecord(ws []Word) (Record, int, error) {
	if len(ws) < 1 {
		return Record{}, 0, errors.New("geul/core: record requires a TAG WORD")
	}
	tag := ws[0]
	if tag>>13 != 0 {
		return Record{}, 0, fmt.Errorf("geul/core: TAG WORD 0x%04X prefix is not 000", uint16(tag))
	}
	k, codec := tagCodec(tag)
	r := Record{Kind: k, Codec: codec}
	total, head, err := decodeLength(ws, uint8(tag&0xF))
	if err != nil {
		return Record{}, 0, err
	}
	if total < head || len(ws) < total {
		return Record{}, 0, errors.New("geul/core: record Length runs past the input")
	}
	body := ws[head:total]
	w := IssuerWidth(r.Kind, r.Codec)
	if w > 0 && len(body) < w+FormatWords {
		return Record{}, 0, errors.New("geul/core: ISSUER extension truncated")
	}
	if w > 0 {
		r.Issuer = append([]Word(nil), body[:w]...)
		r.Format = uint32(body[w])<<16 | uint32(body[w+1])
		body = body[w+FormatWords:]
	}
	r.Payload = body
	return r, total, nil
}
