//ff:func feature=core type=encoder control=sequence
//ff:what ContentHash — SHA-256 of a document from DOC-* to just before END, validation-class Meta left out
package core

import "crypto/sha256"

// ContentHash computes a document's content hash (GEUL t002 OVERVIEW 계보): the
// SHA-256 of its words from the opening DOC-* record up to, not including, END.
// Validation-class Meta records (Meta C >= 0x25, IsValidationMeta: ISSUER-KEY,
// SIGNATURE, VERDICT, …) are left out, and the DOC-* record is rewritten for the
// remaining length at its narrowest width, so adding or removing signatures does
// not change the hash. ws must be exactly one well-formed document.
func ContentHash(ws []Word) ([32]byte, error) {
	if _, err := SplitDocument(ws); err != nil {
		return [32]byte{}, err
	}
	_, n, _ := DecodeRecord(ws)
	doc := WrapDocument(stripValidation(ws[n:]))
	return sha256.Sum256(PutWords(doc[:len(doc)-1])), nil
}
