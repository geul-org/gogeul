//ff:func feature=pack type=helper control=sequence
//ff:what headIssuer — the Issuer of a Meta DOC-ID record (DOC-POINTER or CODEC-POINTER); false for any other record
package pack

import "github.com/geul-org/gogeul/pkg/core"

// headIssuer returns the Issuer r names when r is Meta DOC-ID: a document's
// DOC-POINTER or a codec document's CODEC-POINTER.
func headIssuer(r core.Record) ([]core.Word, bool) {
	if is, _, err := core.DecodeDocID(r); err == nil {
		return is, true
	}
	is, _, err := core.DecodeDocIDCodec(r)
	return is, err == nil
}
