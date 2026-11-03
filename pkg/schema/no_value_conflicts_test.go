//ff:func feature=schema type=rule control=sequence
//ff:what TestNoValueConflicts — NO-VALUE beside a value or UNKNOWN-VALUE of the same S·P·rank and qualifiers warns E6; another rank or a telling qualifier does not
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestNoValueConflicts(t *testing.T) {
	r := func(ws ...core.Word) rec {
		x, _, err := core.DecodeRecord(ws)
		if err != nil {
			t.Fatal(err)
		}
		return recOf(x)
	}
	none := r(0x08A8, 0x0051, 1, 0x0191, 1304, 0x0141, 30, core.NoValue, 0)
	fed := r(0x08A9, 0x0051, 2, 0x0191, 1304, 0x0141, 30, 0x0141, 53536, 0)
	unknown := r(0x08A8, 0x0051, 3, 0x0191, 1304, 0x0141, 30, core.UnknownValue, 0)
	fedPreferred := r(0x08A9, 0x0051, 2, 0x0191, 1304, 0x0141, 30, 0x0141, 53536, 1)
	from := r(0x08B9, 0x0051, 4, 0x0191, 580, 0x0051, 1, 0x04E2, 0x3138, 0x3131)
	fs := noValueConflicts([]rec{none, fed})
	if len(fs) != 1 || fs[0].Rule != "E6" || fs[0].Severity != Warn {
		t.Fatalf("value beside NO-VALUE: %+v", fs)
	}
	if got := rulesOf(noValueConflicts([]rec{unknown, none})); got != "E6" {
		t.Fatalf("UNKNOWN-VALUE beside NO-VALUE: %q", got)
	}
	if fs := noValueConflicts([]rec{none, fedPreferred}); len(fs) != 0 {
		t.Fatalf("other rank: %+v", fs)
	}
	if fs := noValueConflicts([]rec{none, fed, from}); len(fs) != 0 {
		t.Fatalf("qualifier tells them apart: %+v", fs)
	}
	if fs := noValueConflicts([]rec{none, recOf(core.EncodeLocalRef(1))}); len(fs) != 0 {
		t.Fatalf("NO-VALUE alone: %+v", fs)
	}
}
