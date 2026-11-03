//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckDocOpen — not DOC, wrong payload width, wrong total, too wide, ok
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckDocOpen(t *testing.T) {
	doc := func(c uint16, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: c, Payload: p})
	}
	if fs := checkDocOpen(doc(core.MetaDoc1Word, 10), 10); fs != nil {
		t.Fatalf("ok: %+v", fs)
	}
	for i, c := range []struct {
		r rec
		n int
	}{
		{recOf(core.EncodeUintValue(1)), 2},
		{doc(core.MetaDoc1Word, 0, 10), 10},
		{doc(core.MetaDoc1Word, 9), 10},
		{doc(core.MetaDoc2Word, 0, 10), 10},
	} {
		if fs := checkDocOpen(c.r, c.n); rulesOf(fs) != "D1" {
			t.Fatalf("case %d: %+v", i, fs)
		}
	}
}
