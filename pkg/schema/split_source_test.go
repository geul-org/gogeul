//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what TestSplitSource — SOURCE cut into Target, Origin and the rest; a broken Edge Id, Target or Origin is false
package schema

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSplitSource(t *testing.T) {
	edge := func(ws ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: ws})
	}
	target, origin, rest, ok := splitSource(edge(0x0051, 9, 0x0051, 2, 0x0141, 0x8EE2, 7))
	if !ok || !slices.Equal(target, []core.Word{0x0051, 2}) || !slices.Equal(origin, []core.Word{0x0141, 0x8EE2}) || !slices.Equal(rest, []core.Word{7}) {
		t.Fatalf("%v %v %v %v", target, origin, rest, ok)
	}
	for i, x := range []rec{
		edge(0x0089),
		edge(0x0051, 9, 0x0089),
		edge(0x0051, 9, 0x0051, 2, 0x00A1),
	} {
		if _, _, _, ok := splitSource(x); ok {
			t.Fatalf("case %d: want false", i)
		}
	}
}
