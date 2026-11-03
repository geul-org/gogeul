//ff:func feature=schema type=rule control=sequence
//ff:what TestHeadEdgeRec — a document SOURCE is noted and lets its QUALIFIER count as a HEAD Edge; other records are not
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestHeadEdgeRec(t *testing.T) {
	st := &placeState{}
	qual := recOf(edgeOf(core.EdgeQualifier, 4, 0x0191, 813, 0x0051, 3, 0x0141, 42))
	if headEdgeRec(st, qual) {
		t.Fatal("QUALIFIER before its SOURCE")
	}
	if !headEdgeRec(st, recOf(edgeOf(core.EdgeSource, 3, 0x0051, 0, 0x0141, 42))) || !st.sources[3] {
		t.Fatalf("document SOURCE: %+v", st)
	}
	if !headEdgeRec(st, qual) || st.sources[4] {
		t.Fatalf("its QUALIFIER: %+v", st)
	}
	if headEdgeRec(st, recOf(edgeOf(core.EdgeSource, 5, 0x0051, 4, 0x0141, 42))) {
		t.Fatal("SOURCE on a record")
	}
}
