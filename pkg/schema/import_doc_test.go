//ff:func feature=schema type=helper control=sequence
//ff:what importDoc — test fixture: IMPORT-DOC of Issuer 0x0000's document, version, with the given IMPORT ID words
package schema

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// importDoc builds an IMPORT-DOC of DOC-POINTER(0x00000000, 0xffff…0001, version)
// followed by the IMPORT ID words id.
func importDoc(version core.Word, id ...core.Word) rec {
	ptr := []core.Word{0x0048, 0, 0, 0xffff, 0, 0, 1, 0, version}
	return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: slices.Concat(ptr, id)})
}
