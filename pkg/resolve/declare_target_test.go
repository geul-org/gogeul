//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what TestDeclareTarget — the one DECLARE's target; none, two and an ISSUER target fail
package resolve

import (
	"github.com/geul-org/gogeul/pkg/core"
	"testing"
)

func TestDeclareTarget(t *testing.T) {
	d := core.EncodeDeclare(core.Declare{EdgeID: 1, Target: 0x0B80})
	if w, ok := declareTarget([]core.Record{core.EncodeUintValue(1), d}); !ok || w != 0x0B80 {
		t.Fatalf("got %04X %v", w, ok)
	}
	iss := core.EncodeDeclare(core.Declare{EdgeID: 2, Target: core.TargetWord(core.KindValue, core.CodecValueIssuer2Word), Issuer: []core.Word{0, 2}, Format: 2})
	for i, body := range [][]core.Record{nil, {d, d}, {iss}} {
		if _, ok := declareTarget(body); ok {
			t.Fatalf("bad %d: want false", i)
		}
	}
}
