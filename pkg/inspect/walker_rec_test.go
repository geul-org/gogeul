//ff:func feature=inspect type=decoder control=sequence
//ff:what TestWalkerRec — an embedded record becomes one Record part; a bad record or a bad walk adds nothing
package inspect

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestWalkerRec(t *testing.T) {
	w := &walker{ws: []core.Word{0x0051, 7, 0xFFFF}}
	w.rec("ID")
	if len(w.out) != 1 || !w.out[0].Record || w.pos != 2 || w.bad {
		t.Fatalf("got %+v", w)
	}
	w.rec("bad")
	if !w.bad || len(w.out) != 1 {
		t.Fatalf("bad record: %+v", w)
	}
	w.rec("after")
	if len(w.out) != 1 {
		t.Fatalf("after bad: %+v", w)
	}
}
