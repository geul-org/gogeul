//ff:func feature=pack type=helper control=sequence
//ff:what isHex — whether s is canonical GEUL hex of exactly digits digits
package pack

// isHex reports whether s is `0x` and digits lowercase hex digits (parseHex).
func isHex(s string, digits int) bool {
	_, ok := parseHex(s, digits)
	return ok
}
