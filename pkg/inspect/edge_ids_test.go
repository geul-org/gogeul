//ff:func feature=inspect type=decoder control=sequence
//ff:what TestEdgeIDs — Edges get their leading LOCAL; a typed Edge Id, non-Edges and Edges without a LOCAL keep theirs
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEdgeIDs(t *testing.T) {
	es := []Entry{
		{Record: core.EncodeFieldDecl(core.FieldDecl{EdgeID: 4, Width: 1})},
		{Record: core.EncodeFieldDecl(core.FieldDecl{EdgeID: 5, Width: 1}), ID: 9},
		{Record: core.EncodeUintValue(1)},
		{Record: core.Record{Kind: core.KindEdge, Codec: 0x05, Payload: []core.Word{0x0C22}}},
	}
	edgeIDs(es)
	if es[0].ID != 4 || es[1].ID != 9 || es[2].ID != 0 || es[3].ID != 0 {
		t.Fatalf("got %d %d %d %d", es[0].ID, es[1].ID, es[2].ID, es[3].ID)
	}
}
