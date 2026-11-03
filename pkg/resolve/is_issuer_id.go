//ff:func feature=resolve type=helper control=sequence
//ff:what isIssuerID — whether s is an Issuer ID written as 8, 16 or 64 hex digits (zero-padded forms included)
package resolve

// isIssuerID reports whether s is an Issuer ID as an address writes it: `0x`
// and 8, 16 or 64 hex digits (2, 4 or 16 words). A shorter ID padded
// with leading zeros to a wider form is the same Issuer (normalIssuer).
func isIssuerID(s string) bool {
	_, ok := normalIssuer(s)
	return ok
}
