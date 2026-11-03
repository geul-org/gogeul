//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what docIssuer — the Issuer of a document's DOC-ID (DOC-POINTER or CODEC-POINTER) in its HEAD; false when it has none
package pack

import "github.com/geul-org/gogeul/pkg/core"

// docIssuer reads the HEAD of the GEUL document b and returns the Issuer its
// Meta DOC-ID names, a document's or a codec document's. A document without
// DOC-ID or bytes that do not split give false.
func docIssuer(b []byte) ([]core.Word, bool) {
	ws, err := core.ReadWords(b)
	if err != nil {
		return nil, false
	}
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return nil, false
	}
	for _, r := range doc.Head {
		if is, _, err := core.DecodeDocID(r); err == nil {
			return is, true
		}
		if is, _, err := core.DecodeDocIDCodec(r); err == nil {
			return is, true
		}
	}
	return nil, false
}
