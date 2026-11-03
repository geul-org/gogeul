//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeEdgeDst — FuncRef and REF dsts decode; unknown relations and provenance fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDecodeEdgeDst(t *testing.T) {
	p, pos := EdgePacket{Prop: PropDo}, 0
	if err := decodeEdgeDst([]core.Word{0x0051, 3, 9}, &pos, &p, testWire); err != nil || p.DstFunc == nil || p.DstFunc.Verb != 9 {
		t.Fatalf("func: %v %+v", err, p)
	}
	p, pos = EdgePacket{Prop: PropAttacks}, 0
	if err := decodeEdgeDst([]core.Word{0x0051, 4}, &pos, &p, testWire); err != nil || p.DstTID != 4 {
		t.Fatalf("ref: %v %+v", err, p)
	}
	for _, prop := range []PropCode{0x0FFF, PropProvenance} {
		p, pos = EdgePacket{Prop: prop}, 0
		if err := decodeEdgeDst([]core.Word{1}, &pos, &p, testWire); err == nil {
			t.Fatalf("0x%04X: want error", uint16(prop))
		}
	}
}
