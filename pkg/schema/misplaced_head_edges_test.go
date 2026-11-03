//ff:func feature=schema type=rule control=sequence
//ff:what TestMisplacedHeadEdges — E8 for a document SOURCE or a QUALIFIER on a HEAD SOURCE in the BODY; BODY SOURCEs and QUALIFIERs on BODY Edges pass
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMisplacedHeadEdges(t *testing.T) {
	head := []rec{recOf(edgeOf(core.EdgeSource, 3, 0x0051, 0, 0x0141, 42))}
	body := []rec{
		recOf(edgeOf(core.EdgeSource, 5, 0x0051, 0, 0x0141, 42)),
		recOf(edgeOf(core.EdgeQualifier, 6, 0x0191, 813, 0x0051, 3, 0x0141, 42)),
		recOf(edgeOf(core.EdgeSource, 7, 0x0051, 9, 0x0141, 42)),
		recOf(edgeOf(core.EdgeQualifier, 8, 0x0191, 813, 0x0051, 7, 0x0141, 42)),
	}
	fs := misplacedHeadEdges(docParts{Head: head, Body: body})
	if rulesOf(fs) != "E8,E8" || fs[0].Actual != "BODY의 SOURCE" || fs[1].Actual != "BODY의 QUALIFIER" {
		t.Fatalf("%+v", fs)
	}
}
