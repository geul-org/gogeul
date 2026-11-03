//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestTakeForce — reads an ATTACKS force in [0, 1]; other relations, negatives, >1 and non-decimals fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// TestTakeForce reads 0.5 and 0 (C=0, E=0) on ATTACKS and checks that pos moves
// past the Value, then rejects a force on SUPERSEDES, -0.5, 2 and a non-Value.
func TestTakeForce(t *testing.T) {
	pos := 0
	if v, err := takeForce([]core.Word{0x0483, 0xFFFF, 0xFFFF, 5, 0x0040}, &pos, PropAttacks); err != nil || *v != 5000 || pos != 4 {
		t.Fatalf("0.5: %v %v pos %d", v, err, pos)
	}
	pos = 0
	if v, err := takeForce([]core.Word{0x0483, 0, 0, 0}, &pos, PropAttacks); err != nil || *v != 0 {
		t.Fatalf("0: %v %v", v, err)
	}
	bad := []struct {
		ws   []core.Word
		prop PropCode
	}{
		{[]core.Word{0x0483, 0xFFFF, 0xFFFF, 5}, PropSupersedes},
		{[]core.Word{0x0483, 0xFFFF, 0xFFFF, 0xFFFB}, PropAttacks},
		{[]core.Word{0x0483, 0, 0, 2}, PropAttacks},
		{[]core.Word{0x0800}, PropAttacks},
	}
	for i, c := range bad {
		pos = 0
		if _, err := takeForce(c.ws, &pos, c.prop); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
