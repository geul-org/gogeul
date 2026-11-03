//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestEncodeEdge — DO-ONCE and ATTACKS-with-force goldens plus every EncodeEdge error
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestEncodeEdge pins a `do bank.withdraw once` LINK (Edge Codec 0x72, TAG 0x112L, flag word 0)
// and checks the errors: provenance, unknown relation, a FuncRef relation
// without DstFunc, a gate flag without a Gate, a bad gate, a bad force and a
// SPAN overflow. An ATTACKS at force 0.5 is pinned word for word (the example of
// the geul.org LINK page).
func TestEncodeEdge(t *testing.T) {
	got, err := EncodeEdge(EdgePacket{ID: 5, Prop: PropDo, Src: 2, DstFunc: &FuncRef{Alias: 3, Verb: 0xC200}, Flags: EdgeFlagOnce}, testWire)
	if want := []core.Word{0x112A, 0x0051, 5, 0x0051, 2, 0x0051, 3, 0xC200, 0, 0, 0}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("do once: %04X %v want %04X", got, err, want)
	}
	half := VFix(5000)
	got, err = EncodeEdge(EdgePacket{ID: 5, Prop: PropAttacks, Src: 3, DstTID: 2, Flags: EdgeFlagForce, Force: &half, Span: Span{Off: 0x40, Len: 0x1A}}, testWire)
	if want := []core.Word{0x108D, 0x0051, 5, 0x0051, 3, 0x0051, 2, 0x0010, 0x0683, 0xFFFF, 0xFFFF, 5, 0x0040, 0x001A}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("attacks force: %04X %v want %04X", got, err, want)
	}
	bad := []EdgePacket{
		{Prop: PropProvenance},
		{Prop: 0x0FFF},
		{Prop: PropDo},
		{Prop: PropAttacks, Flags: EdgeFlagGate},
		{Prop: PropAttacks, Flags: EdgeFlagGate, Gate: &Gate{Percent: 101}},
		{Prop: PropAttacks, Flags: EdgeFlagForce},
		{Prop: PropAttacks, Span: Span{Len: 0x10000}},
	}
	for i, p := range bad {
		if _, err := EncodeEdge(p, testWire); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
