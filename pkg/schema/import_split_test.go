//ff:func feature=schema type=decoder control=sequence
//ff:what TestImportSplit — the words after a DOC-POINTER; another Reference or a broken record fails
package schema

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestImportSplit(t *testing.T) {
	if id, ok := importSplit(importDoc(1, 2, 3)); !ok || !slices.Equal(id, []core.Word{2, 3}) {
		t.Fatalf("got %v %v", id, ok)
	}
	if _, ok := importSplit(recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x0051, 1, 2}})); ok {
		t.Fatal("LOCAL accepted")
	}
	if _, ok := importSplit(recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x03D8, 0}})); ok {
		t.Fatal("short accepted")
	}
}
