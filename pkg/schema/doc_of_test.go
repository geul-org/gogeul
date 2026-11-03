//ff:func feature=schema type=helper control=iteration dimension=1
//ff:what docOf — test fixture: a document from HEAD records and BODY records
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// docOf wraps head records, BODY, body records and END in a DOC record. An
// empty head writes no BODY (the canonical form, D3).
func docOf(head, body []core.Record) []byte {
	var in []core.Word
	for _, r := range head {
		in = append(in, enc(r)...)
	}
	if len(head) > 0 {
		in = append(in, 0x1C30)
	}
	for _, r := range body {
		in = append(in, enc(r)...)
	}
	return core.PutWords(core.WrapDocument(append(in, 0x1C40)))
}
