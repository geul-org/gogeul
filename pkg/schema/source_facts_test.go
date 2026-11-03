//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestSourceFacts — E8: good SOURCEs pass; bad layout, a Target that is no Edge Id, an Origin that is a Value, NO-VALUE or a reserved marker fail (ABSENT and UNKNOWN-VALUE pass)
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSourceFacts(t *testing.T) {
	edge := func(ws ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindEdge, Codec: core.EdgeSource, Payload: ws})
	}
	ids := map[uint64]bool{2: true}
	cases := map[string]rec{
		"":      edge(0x0051, 9, 0x0051, 2, 0x0141, 0x8EE2),
		"E8":    edge(0x0051, 9, 0x0051, 3, 0x0051, 20),
		"E8,E8": edge(0x0051, 9, 0x0141, 2, core.NoValue),
	}
	for want, r := range cases {
		if got := rulesOf(sourceFacts(r, ids)); got != want {
			t.Fatalf("want %q: got %q", want, got)
		}
	}
	for i, r := range []rec{
		edge(0x0051, 9, 0x0051, 0, core.UnknownValue),
		edge(0x0051, 9, 0x0051, 2, core.Absent),
		edge(0x0051, 9, 0x0051, 2, 0x0051, 20),
	} {
		if fs := sourceFacts(r, ids); len(fs) != 0 {
			t.Fatalf("good %d: %+v", i, fs)
		}
	}
	for i, r := range []rec{
		edge(0x0051, 9, 0x0051, 2, 0x0141, 5, 0),
		edge(0x0051, 9, 0x0051, 2),
		edge(0x0051, 9, 0x0051, 2, 0x0003),
		edge(0x0051, 9, 0x0051, 2, 0x0621, 5),
	} {
		if got := rulesOf(sourceFacts(r, ids)); got != "E8" {
			t.Fatalf("bad %d: %q", i, got)
		}
	}
}
