//ff:func feature=core type=encoder control=sequence
//ff:what EncodeRecord — serialize a Record as TAG WORD, Length extension, ISSUER extension and Payload
package core

import (
	"errors"
	"fmt"
)

// EncodeRecord serializes r: the TAG WORD, the canonical Length extension, the
// ISSUER extension (Issuer + 2-word codec number) when Codec is an ISSUER code,
// then the Payload. A Codec beyond its Kind's zone, an Issuer whose width does
// not match the codec, or a codec number on a direct codec is a hard error.
func EncodeRecord(r Record) ([]Word, error) {
	if r.Kind > KindMeta || r.Codec > CodecMax(r.Kind) {
		return nil, fmt.Errorf("geul/core: codec 0x%03X is beyond the %v zone", r.Codec, r.Kind)
	}
	w := IssuerWidth(r.Kind, r.Codec)
	if w != len(r.Issuer) {
		return nil, fmt.Errorf("geul/core: codec 0x%03X needs a %d-word Issuer, got %d", r.Codec, w, len(r.Issuer))
	}
	if w == 0 && r.Format != 0 {
		return nil, errors.New("geul/core: codec number on a direct codec")
	}
	body := append([]Word(nil), r.Issuer...)
	if w > 0 {
		body = append(body, Word(r.Format>>16), Word(r.Format))
	}
	body = append(body, r.Payload...)
	return prependTag(r.Kind, r.Codec, len(body), body)
}
