//ff:func feature=pack type=helper control=sequence
//ff:what hashState — the marshaled state of a running SHA-256 (nil for none)
package pack

import (
	"encoding"
	"hash"
)

// hashState saves h (a crypto/sha256 hash, which marshals without error) so
// a later run can compare against it or go on from it.
func hashState(h hash.Hash) []byte {
	if h == nil {
		return nil
	}
	b, _ := h.(encoding.BinaryMarshaler).MarshalBinary()
	return b
}
