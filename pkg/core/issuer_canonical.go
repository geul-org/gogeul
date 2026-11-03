//ff:func feature=core type=helper control=selection
//ff:what IssuerCanonical — whether an Issuer ID is 2, 4 or 16 words at its narrowest width
package core

import "slices"

// IssuerCanonical reports whether issuer is an Issuer ID in canonical form: 2,
// 4 or 16 words, and the shortest of those that holds its value. An Issuer ID
// is its integer value, so a 2-word ID written with leading zero words as 4 or
// 16 words is the same Issuer (IssuerNormal); that is why no 4-word Issuer has
// its top 2 words zero (it starts at 2^32) and no 16-word Issuer its top 12
// words zero (it starts at 2^64). Every 2-word value is canonical; 1 word is
// not an Issuer width.
func IssuerCanonical(issuer []Word) bool {
	switch len(issuer) {
	case 2:
		return true
	case 4:
		return issuer[0]|issuer[1] != 0
	case 16:
		return slices.ContainsFunc(issuer[:12], func(w Word) bool { return w != 0 })
	default:
		return false
	}
}
