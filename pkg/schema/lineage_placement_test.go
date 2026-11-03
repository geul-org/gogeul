//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestLineagePlacement — NEXT-LOCAL-ID in a version list and a diff with DOC-ID and Edges pass; NEXT-LOCAL-ID in a document, DELTA-BASE without DOC-ID or over a non-Edge BODY, DELTA-REMOVE without DELTA-BASE fail M7
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLineagePlacement(t *testing.T) {
	id, _ := core.EncodeDocID([]core.Word{0, 2}, []core.Word{0, 0, 0, 5})
	next, base := recOf(core.EncodeNextLocalID(4)), recOf(core.EncodeDeltaBase(1))
	remove := recOf(core.EncodeDeltaRemove([]uint64{3}))
	pub := recOf(core.EncodePublication(core.Publication{Codec: core.EdgePublication, EdgeID: 1, Issuer: []core.Word{0, 2}, Doc: []core.Word{0, 0, 0, 5}, Version: 1}))
	edge, value := recOf(edgeOf(core.EdgeBind, 2, 0x0051, 1, 0x0721, 0x6100)), recOf(core.EncodeUintValue(1))
	cases := []struct {
		parts docParts
		want  string
	}{
		{docParts{Head: []rec{next}, Body: []rec{pub}}, ""},
		{docParts{Head: []rec{recOf(id), base, remove}, Body: []rec{edge}}, ""},
		{docParts{Head: []rec{recOf(id), next}, Body: []rec{edge}}, "M7"},
		{docParts{Head: []rec{base}, Body: []rec{edge}}, "M7"},
		{docParts{Head: []rec{recOf(id), base}, Body: []rec{edge, value}}, "M7"},
		{docParts{Head: []rec{recOf(id), remove}, Body: []rec{edge}}, "M7"},
		{docParts{Head: []rec{base, remove}, Body: []rec{value}}, "M7,M7"},
	}
	for i, c := range cases {
		if got := rulesOf(lineagePlacement(c.parts)); got != c.want {
			t.Fatalf("case %d: %q want %q", i, got, c.want)
		}
	}
}
