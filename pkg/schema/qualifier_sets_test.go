//ff:func feature=schema type=helper control=sequence
//ff:what TestQualifierSets — QUALIFIERs grouped by Target, order-free; TRIPLEs and a non-LOCAL Target are skipped
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestQualifierSets(t *testing.T) {
	st := func(ws ...core.Word) statement {
		x, _, _ := core.DecodeRecord(ws)
		s, ok := splitStatement(recOf(x))
		if !ok {
			t.Fatalf("%04X", ws)
		}
		return s
	}
	a := st(0x0AB7, 0x0051, 8, 0x0191, 582, 0x0051, 7, core.UnknownValue)
	b := st(0x0AB7, 0x0051, 9, 0x0191, 580, 0x0051, 7, core.NoValue)
	tr := st(0x0AA8, 0x0051, 7, 0x0191, 40, 0x0141, 42, core.NoValue, 0)
	bad := st(0x0AB7, 0x0051, 10, 0x0191, 582, 0x0141, 7, core.UnknownValue)
	m1, m2 := qualifierSets([]statement{a, b, tr, bad}), qualifierSets([]statement{b, a})
	if len(m1) != 1 || m1[7] == "" || m1[7] != m2[7] {
		t.Fatalf("%v %v", m1, m2)
	}
}
