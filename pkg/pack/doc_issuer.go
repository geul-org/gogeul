//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what docIssuer — the Issuer a document belongs to: its DOC-ID's (DOC-POINTER or CODEC-POINTER), else a version list's first record's pointer; false when neither
package pack

import "github.com/geul-org/gogeul/pkg/core"

// docIssuer reads the GEUL document b and returns the Issuer it belongs to:
// the one its Meta DOC-ID names (a document's or a codec document's), or for a
// version list, which has no DOC-ID, the Issuer of its first record's pointer
// (PUBLICATION, DEPRECATION or TOMBSTONE). Bytes that do not split, or a
// document with neither, give false.
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
		if is, ok := headIssuer(r); ok {
			return is, true
		}
	}
	if len(doc.Body) == 0 {
		return nil, false
	}
	p, err := core.DecodePublication(doc.Body[0])
	return p.Issuer, err == nil
}
