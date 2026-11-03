//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestStatementLayout — good TRIPLE and QUALIFIER pass; rank 5, no Rank, two Rank WORDs, a Value Predicate, a non-LOCAL Target, an Edge Object, words after the Value and an unsplittable payload fail E10; other Edges and a missing Edge Id are not E10's
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestStatementLayout(t *testing.T) {
	p, q, v := []core.Word{0x0191, 31}, []core.Word{0x0141, 42}, []core.Word{0x0621, 7}
	local := enc(core.EncodeLocalRef(1))
	cat := func(parts ...[]core.Word) []core.Word {
		var out []core.Word
		for _, x := range parts {
			out = append(out, x...)
		}
		return out
	}
	edge := []core.Word{0x0AA0}
	cases := []struct {
		r    core.Record
		want string
	}{
		{edgeOf(core.EdgeTriple, 1, cat(p, q, v, []core.Word{1})...), ""},
		{edgeOf(core.EdgeTriple, 1, cat(p, q, []core.Word{0x0010}, []core.Word{2})...), ""},
		{edgeOf(core.EdgeTriple, 1, cat(p, q, v, []core.Word{5})...), "E10"},
		{edgeOf(core.EdgeTriple, 1, cat(p, q, v)...), "E10"},
		{edgeOf(core.EdgeTriple, 1, cat(p, q, v, []core.Word{0, 0})...), "E10"},
		{edgeOf(core.EdgeTriple, 1, cat(v, q, v, []core.Word{0})...), "E10"},
		{edgeOf(core.EdgeTriple, 1, cat(p, q, edge, []core.Word{0})...), "E10"},
		{edgeOf(core.EdgeTriple, 1, cat(p, []core.Word{0x0613})...), "E10"},
		{edgeOf(core.EdgeQualifier, 2, cat(p, local, v)...), ""},
		{edgeOf(core.EdgeQualifier, 2, cat(p, q, v)...), "E10"},
		{edgeOf(core.EdgeQualifier, 2, cat(p, local, v, []core.Word{0})...), "E10"},
		{core.Record{Kind: core.KindEdge, Codec: core.EdgeTriple}, ""},
		{edgeOf(core.EdgeBind, 1, 0), ""},
		{core.EncodeUintValue(1), ""},
	}
	for i, c := range cases {
		if got := rulesOf(statementLayout(recOf(c.r))); got != c.want {
			t.Fatalf("case %d: %q want %q", i, got, c.want)
		}
	}
}
