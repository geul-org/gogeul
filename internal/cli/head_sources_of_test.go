//ff:func feature=cli type=helper control=sequence
//ff:what TestHeadSourcesOf — the Edge Ids of document SOURCEs; Meta, QUALIFIERs and SOURCEs on a record are left out
package cli

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/notation"
)

func TestHeadSourcesOf(t *testing.T) {
	zero, _ := core.EncodeRecord(core.EncodeLocalRef(0))
	two, _ := core.EncodeRecord(core.EncodeLocalRef(2))
	q1 := []core.Word{0x0141, 1}
	head := []core.Record{
		core.EncodeSpecVersion(core.SpecVersion),
		edgeRecord(2, notation.Edge{Codec: core.EdgeSource, Rest: append(append([]core.Word{}, zero...), q1...)}),
		edgeRecord(3, notation.Edge{Codec: core.EdgeQualifier, Rest: append(append([]core.Word{0x0191, 813}, two...), q1...)}),
		edgeRecord(4, notation.Edge{Codec: core.EdgeSource, Rest: append(append([]core.Word{}, two...), q1...)}),
	}
	if got := headSourcesOf(head); len(got) != 1 || !got[2] {
		t.Fatalf("%v", got)
	}
}
