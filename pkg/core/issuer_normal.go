//ff:func feature=core type=helper control=sequence
//ff:what IssuerNormal — an Issuer ID of 2/4/16 words reduced to its canonical width (leading zero words dropped)
package core

// IssuerNormal returns the canonical form of an Issuer ID written as 2, 4 or 16
// words: a 16-word ID whose top 12 words are zero is the same Issuer as its low
// 4 words, and a 4-word ID whose top 2 words are zero the same as its low 2
// words. The bool is false for a width other than 2, 4 or 16.
func IssuerNormal(issuer []Word) ([]Word, bool) {
	if len(issuer) != 2 && len(issuer) != 4 && len(issuer) != 16 {
		return nil, false
	}
	if len(issuer) == 16 && !IssuerCanonical(issuer) {
		issuer = issuer[12:]
	}
	if len(issuer) == 4 && !IssuerCanonical(issuer) {
		issuer = issuer[2:]
	}
	return issuer, true
}
