//ff:func feature=pack type=helper control=sequence
//ff:what storedTail — whether a file name is a stored one: versions.geul or a version ID (8 hex digits, not 0).geul
package pack

import (
	"encoding/binary"
	"strings"
)

// storedTail reports whether name is versions.geul or `<버전 ID>.geul` (8
// hex digits, not 0). latest.geul is not stored.
func storedTail(name string) bool {
	if name == listFile {
		return true
	}
	v, ok := parseHex(strings.TrimSuffix(name, ".geul"), 8)
	return ok && strings.HasSuffix(name, ".geul") && binary.BigEndian.Uint32(v) != 0
}
