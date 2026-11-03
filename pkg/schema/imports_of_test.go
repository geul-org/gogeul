//ff:func feature=schema type=rule control=sequence
//ff:what TestImportsOf — IMPORT IDs collected, duplicates reported, other Meta skipped
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestImportsOf(t *testing.T) {
	imp := func(id core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x0048, 0, 2, 0, 0, 0, 2, 0, 0, id}})
	}
	other := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: make([]core.Word, 5)})
	ids, fs := importsOf([]rec{imp(1), other, imp(2), imp(1)})
	if len(ids) != 2 || rulesOf(fs) != "M3" {
		t.Fatalf("got %v %+v", ids, fs)
	}
}
