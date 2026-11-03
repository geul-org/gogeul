//ff:func feature=schema type=helper control=sequence
//ff:what subjectOf — test fixture: DOC-POINTER Reference to Issuer 0x0000, a document and a version
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
)

func subjectOf(doc, version core.Word) core.Record {
	return core.Record{Kind: core.KindReference, Codec: core.CodecDocPointer, Payload: []core.Word{0, 0, 0xffff, 0, 0, doc, 0, version}}
}
