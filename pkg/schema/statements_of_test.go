//ff:func feature=schema type=decoder control=sequence
//ff:what TestStatementsOf — only TRIPLEs and QUALIFIERs are kept, in order
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestStatementsOf(t *testing.T) {
	tr, _, _ := core.DecodeRecord([]core.Word{0x08A8, 0x0051, 7, 0x0191, 40, 0x0141, 42, core.NoValue, 0})
	q, _, _ := core.DecodeRecord([]core.Word{0x08B7, 0x0051, 8, 0x0191, 582, 0x0051, 7, core.UnknownValue})
	sts := statementsOf([]rec{recOf(tr), recOf(core.EncodeLocalRef(3)), recOf(q)})
	if len(sts) != 2 || sts[0].ID != 7 || sts[1].ID != 8 {
		t.Fatalf("%+v", sts)
	}
}
