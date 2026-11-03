//ff:func feature=toulmin type=decoder control=iteration dimension=1
//ff:what TestDecodeClaimObject — REF, integer and SPAN objects decode; bad values, overflow and unknown kinds fail
package toulmin

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDecodeClaimObject(t *testing.T) {
	var c Claim
	pos := 0
	if err := decodeClaimObject([]core.Word{0x0051, 9}, &pos, ObjTID, testWire, &c); err != nil || c.ObjTID != 9 {
		t.Fatalf("REF: %v %+v", err, c)
	}
	pos = 0
	if err := decodeClaimObject([]core.Word{0x0621, 7}, &pos, ObjNum, testWire, &c); err != nil || c.ObjNum != 7 {
		t.Fatalf("integer: %v %+v", err, c)
	}
	pos = 0
	if err := decodeClaimObject([]core.Word{2, 3}, &pos, ObjSpan, testWire, &c); err != nil || c.ObjSpan != (Span{Off: 2, Len: 3}) {
		t.Fatalf("span: %v %+v", err, c)
	}
	bad := []struct {
		ws   []core.Word
		kind ObjKind
	}{
		{[]core.Word{0x1200}, ObjNum},
		{[]core.Word{0x0A40}, ObjNum},
		{[]core.Word{0x0623, 1, 0, 0}, ObjNum},
		{nil, 3},
	}
	for i, b := range bad {
		p := 0
		if err := decodeClaimObject(b.ws, &p, b.kind, testWire, &c); err == nil {
			t.Fatalf("bad %d: want error", i)
		}
	}
}
