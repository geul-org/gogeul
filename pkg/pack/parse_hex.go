//ff:func feature=pack type=decoder control=sequence
//ff:what parseHex — the bytes of `0x` + exactly digits hex digits (either case on input), false otherwise
package pack

import (
	"encoding/hex"
	"strings"
)

// parseHex decodes GEUL hex as an address is read: `0x` then exactly digits
// hex digits. A-F is written upper case (§1.1 120) but read in either case
// (§1.1 122); the canonical spelling is canonSeg's. `0X` is not hex.
func parseHex(s string, digits int) ([]byte, bool) {
	if len(s) != 2+digits || !strings.HasPrefix(s, "0x") {
		return nil, false
	}
	b, err := hex.DecodeString(s[2:])
	return b, err == nil
}
