//ff:func feature=resolve type=decoder control=sequence
//ff:what storedHexID — bytes of a stored folder name: a `0x` hex ID of the given word width with A-F upper case only
package resolve

import "strings"

// storedHexID decodes a name read from storage (a folder of a directory
// archive). Stored names are written with A-F upper case (§1.1 120); a lower
// case folder is not one of them, as a folder of another width is not.
func storedHexID(s string, words int) ([]byte, bool) {
	b, ok := parseHexID(s, words)
	return b, ok && strings.ToUpper(s[2:]) == s[2:]
}
