//ff:func feature=pack type=helper control=sequence
//ff:what canonicalIssuer — whether an Issuer hex is already in canonical width (normalIssuer gives it back)
package pack

// canonicalIssuer reports whether id is an Issuer ID in its canonical width:
// a zero-padded spelling of a narrower Issuer is not.
func canonicalIssuer(id string) bool {
	n, ok := normalIssuer(id)
	return ok && n == id
}
