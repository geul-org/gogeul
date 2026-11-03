//ff:func feature=page type=helper control=sequence
//ff:what pointerFields — Issuer, an ID of idWords words and the 2-word version ID of a DOC- or CODEC-POINTER payload of n words
package page

// pointerFields splits a DOC-POINTER (idWords 4) or CODEC-POINTER (idWords 2)
// payload of n words: the Issuer is what the ID and the 2-word version ID
// leave. A short payload keeps what is there.
func pointerFields(n int, id string, idWords int) []refField {
	w := max(n-idWords-2, 0)
	d := min(idWords, n-w)
	return []refField{{"Issuer", w}, {id, d}, {"버전 ID", n - w - d}}
}
