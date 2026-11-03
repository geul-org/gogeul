//ff:func feature=pack type=writer control=sequence
//ff:what ownerOf — the Issuer set a path belongs to when an archive is split by Issuer ("" when it is not under an Issuer)
package pack

import "strings"

// ownerOf returns the Issuer of `<Issuer>/…` paths, else "".
func ownerOf(p string) string {
	first, _, nested := strings.Cut(p, "/")
	if is, ok := normalIssuer(first); nested && ok {
		return is
	}
	return ""
}
