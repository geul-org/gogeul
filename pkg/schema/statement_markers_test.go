//ff:func feature=schema type=rule control=sequence
//ff:what TestStatementMarkers — E7 over a BODY: a misplaced marker in a TRIPLE is reported, other records are skipped
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestStatementMarkers(t *testing.T) {
	bad, _, _ := core.DecodeRecord([]core.Word{0x0AA8, 0x0051, 1, 0x0191, 31, 0x0141, 30, core.Absent, 0})
	good, _, _ := core.DecodeRecord([]core.Word{0x0AA8, 0x0051, 2, 0x0191, 31, 0x0141, 30, core.UnknownValue, 0})
	if got := rulesOf(statementMarkers([]rec{recOf(bad), recOf(good), recOf(core.EncodeLocalRef(1))})); got != "E7" {
		t.Fatalf("%q", got)
	}
}
