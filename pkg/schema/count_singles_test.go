//ff:func feature=schema type=rule control=sequence
//ff:what TestCountSingles — a repeated single Meta is reported once per repeat
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCountSingles(t *testing.T) {
	m := func(c uint16) rec { return recOf(core.Record{Kind: core.KindMeta, Codec: c}) }
	if fs := countSingles([]rec{m(core.MetaDocID), m(core.MetaEffectiveFrom), m(0x20), m(0x20), m(core.MetaDocID)}); rulesOf(fs) != "M2" {
		t.Fatalf("got %+v", fs)
	}
}
