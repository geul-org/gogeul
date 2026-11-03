//ff:func feature=schema type=rule control=sequence
//ff:what TestListRecord — Edge Id order, PUBLICATION dispatch, BIND refused (§1.1 102), another document and foreign records
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestListRecord(t *testing.T) {
	st := &listState{doc: "0xffff000000000002", docs: map[string]*docVersions{}}
	if fs := listRecord(st, 0, recOf(edgeOf(core.EdgePublication, 1, append([]core.Word{0x0418, 0, 0, 0xffff, 0, 0, 2, 0, 1}, make([]core.Word, 16)...)...))); fs != nil {
		t.Fatalf("publication: %+v", fs)
	}
	if st.docs["0xffff000000000002"] == nil {
		t.Fatalf("publication not applied: %+v", st.docs)
	}
	if fs := listRecord(st, 1, recOf(bindOf(2, subjectOf(2, 0), "sample"))); rulesOf(fs) != "P1" {
		t.Fatalf("bind: %+v", fs)
	}
	if fs := listRecord(st, 5, recOf(bindOf(2, subjectOf(2, 0), "sample"))); rulesOf(fs) != "P2,P1" {
		t.Fatalf("order: %+v", fs)
	}
	other := recOf(edgeOf(core.EdgeDeprecation, 4, 0x0418, 0, 0, 0xffff, 0, 0, 3, 0, 1))
	if fs := listRecord(st, 3, other); rulesOf(fs) != "P1" {
		t.Fatalf("other document: %+v", fs)
	}
	if fs := listRecord(st, 4, recOf(core.EncodeUintValue(1))); rulesOf(fs) != "P1" {
		t.Fatalf("foreign: %+v", fs)
	}
}
