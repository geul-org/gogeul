//ff:func feature=schema type=rule control=sequence
//ff:what TestCheckEffectiveOrder — FROM before UNTIL passes; equal fails M5; one bound alone passes
package schema

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestCheckEffectiveOrder(t *testing.T) {
	at := func(c uint16, d core.Word) rec {
		return recOf(metaOf(c, core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: []core.Word{0, d}}))
	}
	if checkEffectiveOrder([]rec{at(core.MetaEffectiveFrom, 1), at(core.MetaEffectiveUntil, 2)}) != nil {
		t.Fatal("ordered")
	}
	if rulesOf(checkEffectiveOrder([]rec{at(core.MetaEffectiveFrom, 2), at(core.MetaEffectiveUntil, 2)})) != "M5" {
		t.Fatal("equal")
	}
	if checkEffectiveOrder([]rec{at(core.MetaEffectiveFrom, 2)}) != nil {
		t.Fatal("alone")
	}
}
