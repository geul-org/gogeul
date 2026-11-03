//ff:func feature=core type=encoder control=selection
//ff:what EncodeTargetPointer — a DOC-POINTER for a 4-word document ID, a CODEC-POINTER for a 2-word codec ID
package core

import "fmt"

// EncodeTargetPointer builds the pointer that names one version of a
// published document: DOC-POINTER when id is a 4-word document ID,
// CODEC-POINTER when id is a 2-word codec ID. version is 2 words (0 is the
// latest).
func EncodeTargetPointer(issuer, id, version []Word) (Record, error) {
	switch {
	case len(id) == 2 && len(version) == 2:
		return EncodeCodecPointer(issuer, uint32(id[0])<<16|uint32(id[1]), uint32(version[0])<<16|uint32(version[1]))
	case len(id) == 4:
		return EncodeDocPointer(issuer, id, version)
	default:
		return Record{}, fmt.Errorf("geul/core: a document ID is 4 words and a codec ID 2 words (got %d)", len(id))
	}
}
