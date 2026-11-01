//ff:func feature=pack type=helper control=sequence
//ff:what sumOfState — the SHA-256 digest a saved hash state (hashState) stands for
package pack

import (
	"crypto/sha256"
	"encoding"
)

// sumOfState restores a sha256 state and returns its digest; a state that
// does not unmarshal is ErrCorrupt.
func sumOfState(state []byte) ([]byte, error) {
	h := sha256.New()
	if err := h.(encoding.BinaryUnmarshaler).UnmarshalBinary(state); err != nil {
		return nil, corrupt("checkpoint hash state")
	}
	return h.Sum(nil), nil
}
