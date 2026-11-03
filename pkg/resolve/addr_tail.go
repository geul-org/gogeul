//ff:func feature=resolve type=helper control=selection
//ff:what addrTail — the last segment of an address: latest.geul, versions.geul or a version ID (1 or more, 8 hex digits).geul
package resolve

import (
	"encoding/binary"
	"strings"
)

// addrTail sets a's version from the file name: `latest.geul` (Latest),
// `versions.geul` (Versions) or `<버전 ID>.geul`, 8 lowercase hex digits and
// not 0. Anything else is bad.
func addrTail(a Addr, name string, bad error) (Addr, error) {
	switch name {
	case latestFile:
		return a, nil
	case listFile:
		a.Versions = true
		return a, nil
	}
	b, ok := parseHexID(strings.TrimSuffix(name, geulExt), 2)
	if !ok || !strings.HasSuffix(name, geulExt) || binary.BigEndian.Uint32(b) == 0 {
		return Addr{}, bad
	}
	a.Version = binary.BigEndian.Uint32(b)
	return a, nil
}
