//ff:func feature=resolve type=helper control=sequence
//ff:what normalWords — an Issuer ID of 2/4/16 words in canonical width; another width is kept as given (it names no address)
package resolve

import "github.com/geul-org/gogeul/pkg/core"

// normalWords returns issuer in canonical width (core.IssuerNormal), so a
// zero-padded 4- or 16-word ID reads the same Issuer. A width other than 2, 4
// or 16 is returned as it is; its address does not exist.
func normalWords(issuer []core.Word) []core.Word {
	if n, ok := core.IssuerNormal(issuer); ok {
		return n
	}
	return issuer
}
