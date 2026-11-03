//ff:func feature=pack type=helper control=sequence
//ff:what isIssuer — whether s is an Issuer ID as an address writes it: 8, 16 or 64 hex digits
package pack

// isIssuer reports whether s is an Issuer ID of 8, 16 or 64 lowercase hex
// digits (2, 4 or 16 words). A zero-padded wider spelling is the same Issuer
// (normalIssuer).
func isIssuer(s string) bool {
	_, ok := normalIssuer(s)
	return ok
}
