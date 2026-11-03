//ff:func feature=pack type=decoder control=sequence
//ff:what parseHashHex — 64 lowercase hex digits (a manifest or object name) as 32 bytes
package pack

import (
	"encoding/hex"
	"strings"
)

// parseHashHex decodes an object name.
func parseHashHex(s string) ([32]byte, bool) {
	var h [32]byte
	b, err := hex.DecodeString(s)
	ok := err == nil && len(b) == len(h) && strings.ToLower(s) == s
	copy(h[:], b)
	return h, ok
}
