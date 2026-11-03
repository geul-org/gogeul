//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckDocIDCodec — M2 takes a codec document's CODEC-POINTER DOC-ID; IMPORT-DOC takes CODEC-POINTER too
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCheckDocIDCodec(t *testing.T) {
	id, _ := core.EncodeDocIDCodec([]core.Word{0, 0}, 0x0980)
	if fs := checkDocID(recOf(id)); fs != nil {
		t.Fatalf("codec DOC-ID: %+v", fs)
	}
	ver := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: []core.Word{0x03C6, 0, 0, 0, 0x0980, 0, 1}})
	if rulesOf(checkDocID(ver)) != "M2" {
		t.Fatal("codec DOC-ID with a version: want M2")
	}
	imp := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x03C6, 0, 0, 0, 0x0980, 0, 1, 3}})
	if fs := checkImportDoc(imp); fs != nil {
		t.Fatalf("IMPORT-DOC of a codec: %+v", fs)
	}
}
