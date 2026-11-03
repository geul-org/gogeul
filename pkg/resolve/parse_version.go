//ff:func feature=resolve type=resolver control=sequence
//ff:what parseVersion — a version ID written in decimal or as canonical 2-word hex
package resolve

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

// parseVersion reads s as a decimal version ID or as `0x` + 8 hex
// digits (FormatVersion).
func parseVersion(s string) (uint32, error) {
	if b, ok := parseHexID(s, 2); ok {
		return binary.BigEndian.Uint32(b), nil
	}
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("geul/resolve: version %q is neither decimal nor 0x + 8 hex digits", s)
	}
	return uint32(n), nil
}
