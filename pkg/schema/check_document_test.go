//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckDocument — document-wide rules; version list rules only for a list document
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckDocument(t *testing.T) {
	id := recOf(core.Record{Kind: core.KindMeta, Codec: core.MetaDocID, Payload: []core.Word{0, 0, 0, 0, 5}})
	body := []rec{recOf(edgeOf(0, 1)), recOf(edgeOf(0, 1))}
	if fs := checkDocument(docParts{Head: []rec{id}, Body: body}); rulesOf(fs) != "E2" {
		t.Fatalf("plain: %+v", fs)
	}
	head := []rec{id, recOf(edgeOf(core.EdgeSource, 1, 0x0051, 0, 0x0141, 42))}
	if fs := checkDocument(docParts{Head: head, Body: body[:1]}); rulesOf(fs) != "E2" {
		t.Fatalf("Edge Ids are one numbering over HEAD and BODY: %+v", fs)
	}
	list := []rec{recOf(edgeOf(core.EdgePublication, 1, append([]core.Word{0x0418, 0, 0, 0xffff, 0, 0, 2, 0, 1}, make([]core.Word, 16)...)...)), recOf(edgeOf(0, 1))}
	if fs := checkDocument(docParts{Body: list}); rulesOf(fs) != "E2,P2,P1" {
		t.Fatalf("list: %+v", fs)
	}
}
