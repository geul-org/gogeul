//ff:func feature=delta type=encoder control=sequence
//ff:what encodeDoc — a document from HEAD and BODY records, wrapped in its DOC-* record (core.EncodeDocument: no BODY marker for an empty HEAD)
package delta

import "github.com/geul-org/gogeul/pkg/core"

// encodeDoc writes HEAD, BODY marker, BODY and END, wrapped in the narrowest
// DOC-* record, through core.EncodeDocument, which omits the BODY marker when
// the HEAD is empty (the one encoding of that content). The records come from
// decoded documents, so they encode.
func encodeDoc(head, body []core.Record) []byte {
	b, _ := core.EncodeDocument(head, body)
	return b
}
