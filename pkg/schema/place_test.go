//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestPlace — every placement case of D2~D4, HEAD SOURCE and QUALIFIER included
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestPlace(t *testing.T) {
	meta := func(c uint16) rec { return rec{Total: 1, R: core.Record{Kind: core.KindMeta, Codec: c}} }
	val := recOf(core.EncodeUintValue(1))
	st, parts := &placeState{}, &docParts{}
	if fs := place(st, parts, val); fs != nil || !st.inBody || len(parts.Body) != 1 {
		t.Fatalf("implicit body: %+v", fs)
	}
	st, parts = &placeState{}, &docParts{}
	if got := rulesOf(place(st, parts, meta(core.MetaBody))); got != "D3" || !st.inBody {
		t.Fatalf("BODY after an empty HEAD: %q", got)
	}
	st, parts = &placeState{}, &docParts{}
	steps := []struct {
		r    rec
		want string
	}{
		{meta(core.MetaSpecVersion), ""},
		{val, "D3"},
		{meta(core.MetaBody), ""},
		{meta(core.MetaBody), "D3"},
		{meta(core.MetaSpecVersion), "D3"},
		{meta(core.MetaDoc1Word), "D4"},
		{val, ""},
		{meta(core.MetaEnd), ""},
		{val, "D2"},
	}
	for i, s := range steps {
		if got := rulesOf(place(st, parts, s.r)); got != s.want {
			t.Fatalf("step %d: %q", i, got)
		}
	}
	if len(parts.Head) != 1 || len(parts.Body) != 3 {
		t.Fatalf("parts: %d %d", len(parts.Head), len(parts.Body))
	}
	src := recOf(edgeOf(core.EdgeSource, 3, 0x0051, 0, 0x0141, 42))
	qual := recOf(edgeOf(core.EdgeQualifier, 4, 0x0191, 813, 0x0051, 3, 0x0141, 42))
	st, parts = &placeState{}, &docParts{}
	for i, r := range []rec{src, meta(core.MetaSpecVersion), qual, recOf(edgeOf(core.EdgeSource, 5, 0x0051, 4, 0x0141, 42))} {
		want := map[int]string{3: "D3"}[i]
		if got := rulesOf(place(st, parts, r)); got != want {
			t.Fatalf("HEAD Edge step %d: %q", i, got)
		}
	}
	if len(parts.Head) != 3 || len(parts.Body) != 1 {
		t.Fatalf("HEAD Edges: %d %d", len(parts.Head), len(parts.Body))
	}
}
