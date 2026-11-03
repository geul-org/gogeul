//ff:func feature=cli type=helper control=sequence
//ff:what TestDraftAppendEdge — a document SOURCE and its QUALIFIER go last in the HEAD; a SOURCE on a record and a QUALIFIER on it go to the BODY
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/notation"
)

func TestDraftAppendEdge(t *testing.T) {
	ref := func(id uint64) []core.Word { ws, _ := core.EncodeRecord(core.EncodeLocalRef(id)); return ws }
	q1 := []core.Word{0x0141, 1}
	d := &draft{head: []core.Record{core.EncodeSpecVersion(core.SpecVersion)}}
	d.appendEdge(edgeRecord(2, notation.Edge{Codec: core.EdgeSource, Rest: append(ref(0), q1...)}))
	d.appendEdge(edgeRecord(3, notation.Edge{Codec: core.EdgeQualifier, Rest: append(append([]core.Word{0x0191, 813}, ref(2)...), q1...)}))
	d.appendEdge(edgeRecord(4, notation.Edge{Codec: core.EdgeSource, Rest: append(ref(9), q1...)}))
	d.appendEdge(edgeRecord(5, notation.Edge{Codec: core.EdgeQualifier, Rest: append(append([]core.Word{0x0191, 813}, ref(4)...), q1...)}))
	if len(d.head) != 3 || len(d.body) != 2 || d.head[2].Codec != core.EdgeQualifier {
		t.Fatalf("head %d body %d", len(d.head), len(d.body))
	}
}
