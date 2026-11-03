//ff:func feature=resolve type=decoder control=sequence
//ff:what parseHexID — bytes of a `0x` hex ID of the given word width, the digits read in either case
package resolve

import (
	"encoding/hex"
	"strings"
)

// parseHexID decodes s, which must be `0x` followed by exactly words×4 hex
// digits (GEUL-ORG ID 표기), and reports whether it was. IDs are written with
// A-F upper case (§1.1 120, core.FormatID) and read in either case (§1.1 122).
func parseHexID(s string, words int) ([]byte, bool) {
	if len(s) != 2+words*4 || !strings.HasPrefix(s, "0x") {
		return nil, false
	}
	b, err := hex.DecodeString(s[2:])
	return b, err == nil
}
