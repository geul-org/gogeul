//ff:func feature=schema type=rule control=sequence
//ff:what TestSourceEdges — E8 over HEAD and BODY: a SOURCE on a declared TRIPLE or the document passes, one on a missing Edge or with NO-VALUE fails, other records are skipped
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSourceEdges(t *testing.T) {
	tr, _, _ := core.DecodeRecord([]core.Word{0x08A9, 0x0051, 2, 0x0191, 31, 0x0141, 42, 0x0141, 5, 0})
	good := core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: []core.Word{0x0051, 9, 0x0051, 2, 0x0141, 0x8EE2}}
	bad := core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: []core.Word{0x0051, 10, 0x0051, 4, 0x0141, 0x8EE2}}
	issuer := core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Issuer: []core.Word{1}, Payload: []core.Word{0x0051, 11}}
	body := []rec{recOf(tr), recOf(good), recOf(bad), recOf(issuer), recOf(core.EncodeLocalRef(3))}
	if got := rulesOf(sourceEdges(docParts{Body: body})); got != "E8" {
		t.Fatalf("%q", got)
	}
	doc := core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: []core.Word{0x0051, 12, 0x0051, 0, core.Absent}}
	head := core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: []core.Word{0x0051, 13, 0x0051, 0, core.NoValue}}
	if got := rulesOf(sourceEdges(docParts{Head: []rec{recOf(doc), recOf(head)}, Body: body[:2]})); got != "E8" {
		t.Fatalf("HEAD SOURCEs are checked too: %q", got)
	}
}
