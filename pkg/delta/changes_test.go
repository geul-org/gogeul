//ff:func feature=delta type=decoder control=iteration dimension=1
//ff:what TestChanges — HEAD added and removed, Edges removed, changed and added by Edge Id; a base or next that is not a snapshot fails
package delta

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestChanges(t *testing.T) {
	base := encodeDoc([]core.Record{core.EncodeSpecVersion(2), core.EncodeNextLocalID(4)}, []core.Record{bindOf(1, 9, "a"), bindOf(2, 9, "b"), bindOf(3, 9, "c")})
	next := encodeDoc([]core.Record{core.EncodeSpecVersion(2), core.EncodeNextLocalID(5)}, []core.Record{bindOf(1, 9, "a"), bindOf(3, 9, "C"), bindOf(4, 9, "d")})
	cs, err := Changes(base, next)
	if err != nil {
		t.Fatal(err)
	}
	want := []struct {
		op   byte
		head bool
		id   uint64
	}{{'-', true, 0}, {'+', true, 0}, {'-', false, 2}, {'~', false, 3}, {'+', false, 4}}
	if len(cs) != len(want) {
		t.Fatalf("%d changes: %+v", len(cs), cs)
	}
	for i, w := range want {
		if cs[i].Op != w.op || cs[i].Head != w.head || cs[i].ID != w.id {
			t.Fatalf("%d: %+v", i, cs[i])
		}
	}
	bad := encodeDoc(nil, []core.Record{core.EncodeUintValue(2)})
	if _, err := Changes(bad, next); err == nil {
		t.Fatal("bad base: want error")
	}
	if _, err := Changes(base, bad); err == nil {
		t.Fatal("bad next: want error")
	}
}
