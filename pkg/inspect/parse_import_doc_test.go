//ff:func feature=inspect type=decoder control=iteration dimension=1
//ff:what TestParseImportDoc — latest and pinned versions, a 2-word IMPORT ID; a broken record, no DOC-POINTER, no IMPORT ID and a non-minimal ID fail
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestParseImportDoc(t *testing.T) {
	latest := []core.Word{0x0418, 0, 2, 0, 0, 0, 2, 0, 0, 3}
	d, ok := parseImportDoc(core.Record{Codec: core.MetaImportDoc, Payload: latest})
	if !ok || d.ID != 3 || len(d.Issuer) != 2 || d.Doc[3] != 2 || d.Version != 0 {
		t.Fatalf("latest: %+v %v", d, ok)
	}
	pinned := []core.Word{0x0418, 0, 2, 0, 0, 0, 2, 1, 9, 1, 3}
	if d, ok := parseImportDoc(core.Record{Codec: core.MetaImportDoc, Payload: pinned}); !ok || d.ID != 0x10003 || d.Version != 0x10009 {
		t.Fatalf("pinned: %+v %v", d, ok)
	}
	bad := []core.Record{
		{Codec: core.MetaImportDoc, Payload: []core.Word{0xFFFF}},
		{Codec: core.MetaImportDoc, Payload: []core.Word{3, 0, 2, 0, 0, 0, 2, 0, 0}},
		{Codec: core.MetaImportDoc, Payload: latest[:9]},
		{Codec: core.MetaImportDoc, Payload: append(append([]core.Word{}, latest[:9]...), 0, 3)},
	}
	for i, r := range bad {
		if _, ok := parseImportDoc(r); ok {
			t.Fatalf("bad %d: want not ok", i)
		}
	}
}
