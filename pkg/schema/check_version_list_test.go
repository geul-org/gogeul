//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckVersionList — IMPORT-DOC, a foreign record, Edge Id order and a valid list
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckVersionList(t *testing.T) {
	imp := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaImportDoc, Payload: []core.Word{0x0048, 0, 2, 0, 0, 0, 2, 0, 0, 1}})
	p1 := recOf(edgeOf(core.EdgePublication, 1, append([]core.Word{0x0048, 0, 0, 0xffff, 0, 0, 2, 0, 1}, make([]core.Word, 16)...)...))
	p2 := recOf(edgeOf(core.EdgeDeprecation, 3, 0x0048, 0, 0, 0xffff, 0, 0, 2, 0, 1))
	if fs := checkVersionList(docParts{Body: []rec{p1, p2}}); rulesOf(fs) != "P2" {
		t.Fatalf("order: %+v", fs)
	}
	if fs := checkVersionList(docParts{Head: []rec{imp}, Body: []rec{p1, recOf(core.EncodeUintValue(1))}}); rulesOf(fs) != "P1,P1" {
		t.Fatalf("foreign: %+v", fs)
	}
	if fs := checkVersionList(docParts{Body: []rec{p1}}); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
}
