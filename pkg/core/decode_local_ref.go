//ff:func feature=core type=decoder control=sequence
//ff:what DecodeLocalRef — the ID of a LOCAL Reference record (Quantifier Identified)
package core

import "fmt"

// DecodeLocalRef returns the target ID of r, which must be a LOCAL Reference
// (LOCAL × Identified, 0x005L) whose ID is in minimal words and fits uint64.
func DecodeLocalRef(r Record) (uint64, error) {
	if r.Kind != KindReference || r.Codec != CodecLocal {
		return 0, fmt.Errorf("geul/core: expected LOCAL Reference, got kind %d codec 0x%02X", r.Kind, r.Codec)
	}
	return wordsToUint(r.Payload)
}
