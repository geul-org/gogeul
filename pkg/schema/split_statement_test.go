//ff:func feature=schema type=decoder control=iteration dimension=1
//ff:what TestSplitStatement — TRIPLE and QUALIFIER cut into Edge Id, three records and the rest; other records, no Edge Id or a broken record are false
package schema

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSplitStatement(t *testing.T) {
	r := func(ws ...core.Word) rec {
		x, _, err := core.DecodeRecord(ws)
		if err != nil {
			t.Fatal(err)
		}
		return recOf(x)
	}
	st, ok := splitStatement(r(0x08A8, 0x0051, 7, 0x0191, 40, 0x0141, 42, core.NoValue, 0))
	if !ok || st.ID != 7 || st.Codec != core.EdgeTriple || !slices.Equal(st.Parts[2], []core.Word{core.NoValue}) || !slices.Equal(st.Rest, []core.Word{0}) {
		t.Fatalf("TRIPLE: %+v %v", st, ok)
	}
	st, ok = splitStatement(r(0x08B7, 0x0051, 8, 0x0191, 582, 0x0051, 7, core.UnknownValue))
	if !ok || st.Codec != core.EdgeQualifier || len(st.Rest) != 0 || !slices.Equal(st.Parts[1], []core.Word{0x0051, 7}) {
		t.Fatalf("QUALIFIER: %+v %v", st, ok)
	}
	for i, x := range []rec{
		recOf(core.EncodeLocalRef(1)),
		r(0x08A3, 0x0141, 1, 0),
		r(0x08A5, 0x0051, 7, 0x0191, 40, 0x0061),
		r(0x0847, 0x0051, 7, 0x0191, 40, 0x0141, 42, 0x0000),
	} {
		if _, ok := splitStatement(x); ok {
			t.Fatalf("case %d: want false", i)
		}
	}
}
