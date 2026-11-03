//ff:func feature=core type=decoder control=sequence
//ff:what DecodePublication — a PUBLICATION, DEPRECATION or TOMBSTONE Edge as a Publication
package core

import (
	"errors"
	"fmt"
)

// DecodePublication reads a version list record (GEUL t002 docs/edge/PUBLICATION):
// the LOCAL Edge Id, the DOC-POINTER (or, for a codec, CODEC-POINTER) of the
// version (version ID not 0) and the codec's tail, whose length is fixed per codec.
func DecodePublication(r Record) (Publication, error) {
	tail, ok := publicationTail(r.Codec)
	if r.Kind != KindEdge || !ok {
		return Publication{}, fmt.Errorf("geul/core: not a version list record (kind %d codec 0x%02X)", r.Kind, r.Codec)
	}
	ref, n, err := DecodeRecord(r.Payload)
	if err != nil {
		return Publication{}, err
	}
	id, err := DecodeLocalRef(ref)
	if err != nil {
		return Publication{}, err
	}
	ptr, m, err := DecodeRecord(r.Payload[n:])
	if err != nil {
		return Publication{}, err
	}
	issuer, doc, ver, err := DecodeTargetPointer(ptr)
	if err != nil {
		return Publication{}, err
	}
	rest := r.Payload[n+m:]
	if len(rest) != tail {
		return Publication{}, fmt.Errorf("geul/core: version list record needs %d words after its DOC-POINTER, got %d", tail, len(rest))
	}
	p := Publication{Codec: r.Codec, EdgeID: id, Issuer: issuer, Doc: doc, Version: uint32(ver[0])<<16 | uint32(ver[1])}
	if p.Version == 0 {
		return Publication{}, errors.New("geul/core: version ID 0 in a version list record")
	}
	fillPublicationTail(&p, rest)
	return p, nil
}
