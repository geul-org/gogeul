//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckImportDoc — a missing DOC-POINTER, a missing IMPORT ID and a non-minimal IMPORT ID fail M3
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckImportDoc(t *testing.T) {
	bad := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x0051, 1, 0, 0, 0, 0, 2, 0, 0}})
	if checkImportDoc(importDoc(0, 1)) != nil || rulesOf(checkImportDoc(bad)) != "M3" || rulesOf(checkImportDoc(importDoc(0))) != "M3" || rulesOf(checkImportDoc(importDoc(0, 0, 1))) != "M3" {
		t.Fatal("wrong")
	}
}
