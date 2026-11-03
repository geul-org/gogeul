//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckDocIDCodec — M2 takes a codec document's CODEC-POINTER DOC-ID; IMPORT-DOC and BIND Subjects take CODEC-POINTER too
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
	subj := core.Record{Kind: core.KindReference, Codec: core.CodecCodecPointer, Payload: []core.Word{0, 0, 0, 0x0980, 0, 0}}
	for _, c := range []core.Record{subj} {
		if d, ok := bindDoc(c); !ok || d != "0x00000980" {
			t.Fatalf("BIND codec Subject: %q %v", d, ok)
		}
	}
}
