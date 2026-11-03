//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestMarkerFacts — UNKNOWN and NONE pass as Object/Value; markers in Predicate, Subject, Property or Target, ABSENT or a reserved marker as Object fail E7
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMarkerFacts(t *testing.T) {
	st := func(ws ...core.Word) statement {
		x, _, err := core.DecodeRecord(ws)
		s, ok := splitStatement(recOf(x))
		if err != nil || !ok {
			t.Fatalf("%04X", ws)
		}
		return s
	}
	for i, s := range []statement{
		st(0x08A8, 0x0051, 1, 0x0191, 40, 0x0141, 30, core.NoValue, 0),
		st(0x08B7, 0x0051, 2, 0x0191, 582, 0x0051, 1, core.UnknownValue),
		st(0x08A9, 0x0051, 3, 0x0191, 31, 0x0141, 30, 0x0141, 5, 0),
	} {
		if fs := markerFacts(s); len(fs) != 0 {
			t.Fatalf("ok %d: %+v", i, fs)
		}
	}
	for want, s := range map[string]statement{
		"Predicate": st(0x08A8, 0x0051, 1, core.UnknownValue, 0x0141, 30, 0x0141, 5, 0),
		"Subject":   st(0x08A8, 0x0051, 1, 0x0191, 31, core.NoValue, 0x0141, 5, 0),
		"Object":    st(0x08A8, 0x0051, 1, 0x0191, 31, 0x0141, 30, core.Absent, 0),
		"Value":     st(0x08B8, 0x0051, 2, 0x0191, 582, 0x0051, 1, 0x0011, 7),
		"Property":  st(0x08B7, 0x0051, 2, core.Absent, 0x0051, 1, 0x0141, 5),
	} {
		fs := markerFacts(s)
		if len(fs) != 1 || fs[0].Rule != "E7" || fs[0].Where.Part != want {
			t.Fatalf("%s: %+v", want, fs)
		}
	}
}
