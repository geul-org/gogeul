//ff:func feature=pack type=decoder control=sequence
//ff:what parseLatest — `<generation> <manifest sha256>\n` in its one canonical spelling
package pack

import (
	"fmt"
	"strings"
)

// parseLatest reads latest.txt; any other spelling (spaces, leading zeros,
// upper case, missing newline) is ErrCorrupt.
func parseLatest(b []byte) (uint64, string, error) {
	var gen uint64
	var name string
	_, err := fmt.Sscanf(strings.TrimSuffix(string(b), "\n"), "%d %s", &gen, &name)
	_, hexOK := parseHashHex(name)
	if err != nil || gen == 0 || !hexOK || fmt.Sprintf("%d %s\n", gen, name) != string(b) {
		return 0, "", corrupt("latest.txt")
	}
	return gen, name, nil
}
