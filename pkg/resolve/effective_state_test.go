//ff:func feature=resolve type=rule control=iteration dimension=1
//ff:what TestEffectiveState — not yet, in effect, expired, no bounds; malformed bounds fail
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"testing"
)

func TestEffectiveState(t *testing.T) {
	from, until := uint32(core.MetaEffectiveFrom), uint32(core.MetaEffectiveUntil)
	doc, _ := inspect.Decode(sampleEffectiveDoc(5, "x", [2]uint32{from, 200}, [2]uint32{until, 300}))
	for at, want := range map[int64]effState{199: stateNotYet, 200: stateInEffect, 299: stateInEffect, 300: stateExpired} {
		if st, err := effectiveState(doc, day(at)); err != nil || st != want {
			t.Fatalf("day %d: %v %v", at, st, err)
		}
	}
	plain, _ := inspect.Decode(sampleDoc(1, "x"))
	if st, err := effectiveState(plain, day(0)); err != nil || st != stateInEffect {
		t.Fatalf("no bounds: %v %v", st, err)
	}
	badMeta := func(c uint16) inspect.Document {
		return inspect.Document{Head: []inspect.Entry{{Record: core.Record{Kind: core.KindMeta, Codec: c}}}}
	}
	for _, c := range []uint16{core.MetaEffectiveFrom, core.MetaEffectiveUntil} {
		if _, err := effectiveState(badMeta(c), day(0)); err == nil {
			t.Fatalf("bad 0x%02X: want error", c)
		}
	}
}
