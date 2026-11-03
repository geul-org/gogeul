//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestEncodeForce — 0.5 and an explicit 1 as trimmed SIGNED-DECIMALs; non-ATTACKS, nil and out-of-range fail
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestEncodeForce appends 0.5 (C=5, E=-1) and 1 (C=1, E=0, kept although it is
// the default) after an existing word, then rejects a force on SUPERSEDES, a
// missing force and forces of -0.0001 and 1.0001.
func TestEncodeForce(t *testing.T) {
	half, one := VFix(5000), VFix(10000)
	got, err := encodeForce([]core.Word{7}, EdgePacket{Prop: PropAttacks, Force: &half})
	if want := []core.Word{7, 0x0683, 0xFFFF, 0xFFFF, 5}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("0.5: %04X %v want %04X", got, err, want)
	}
	got, err = encodeForce(nil, EdgePacket{Prop: PropAttacks, Force: &one})
	if want := []core.Word{0x0683, 0, 0, 1}; err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("1: %04X %v want %04X", got, err, want)
	}
	neg, over := VFix(-1), VFix(10001)
	bad := []EdgePacket{
		{Prop: PropSupersedes, Force: &half},
		{Prop: PropAttacks},
		{Prop: PropAttacks, Force: &neg},
		{Prop: PropAttacks, Force: &over},
	}
	for i, p := range bad {
		if _, err := encodeForce(nil, p); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
