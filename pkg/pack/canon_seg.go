//ff:func feature=pack type=helper control=sequence
//ff:what canonSeg — one path segment `0x<hex>` (optionally `.geul`) with A-F upper case, the stored spelling
package pack

import "strings"

// canonSeg writes the hex digits of a `0x…` segment (an ID, or a version or
// content hash file name ending in .geul) upper case, the spelling paths are
// stored and written in (§1.1 120, 122). Other segments are returned as they
// are.
func canonSeg(s string) string {
	digits, ok := strings.CutPrefix(s, "0x")
	if !ok {
		return s
	}
	stem := strings.TrimSuffix(digits, ".geul")
	return "0x" + strings.ToUpper(stem) + digits[len(stem):]
}
