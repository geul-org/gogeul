//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestCheckMeta — each Meta branch; an unassigned code is not reported (§1.1 112)
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckMeta(t *testing.T) {
	raw := func(c uint16, p ...core.Word) rec {
		return recOf(core.Record{Kind: core.KindMeta, Codec: c, Payload: p})
	}
	cases := []struct {
		r    rec
		want string
	}{
		{raw(core.MetaDocID, 1), "M2"},
		{raw(core.MetaImportDoc, 1), "M3"},
		{raw(core.MetaSpecVersion), "M4"},
		{raw(core.MetaProfile), "M4"},
		{raw(core.MetaDefaultLanguage), "M4"},
		{raw(core.MetaEffectiveFrom), "M5"},
		{raw(core.MetaParent), "M6"},
		{raw(core.MetaDeltaBase), "M7"},
		{raw(core.MetaDeltaRemove), "M7"},
		{raw(core.MetaNextLocalID), "M7"},
		{raw(0x2A), ""},
		{raw(0x3F), ""},
		{raw(0x20), ""},
	}
	for i, c := range cases {
		if got := rulesOf(checkMeta(c.r)); got != c.want {
			t.Fatalf("case %d: %q", i, got)
		}
	}
}
