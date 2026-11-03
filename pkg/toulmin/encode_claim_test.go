//ff:func feature=toulmin type=encoder control=iteration dimension=1
//ff:what TestEncodeClaim — each object kind encodes; SPAN object overflow and unknown kinds fail
package toulmin

import (
	"reflect"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEncodeClaim(t *testing.T) {
	cases := []struct {
		c    Claim
		want []core.Word
	}{
		{Claim{Subject: 4, Predicate: 7, ObjTID: 8, Mode: 1}, []core.Word{0x0051, 4, 7, 0x0051, 8, 1}},
		{Claim{Kind: ObjNum, ObjNum: 7}, []core.Word{0x0051, 0, 0, 0x0421, 7, 0}},
		{Claim{Kind: ObjSpan, ObjSpan: Span{Off: 2, Len: 3}}, []core.Word{0x0051, 0, 0, 2, 3, 0}},
	}
	for _, c := range cases {
		if got, err := encodeClaim(nil, c.c, testWire); err != nil || !reflect.DeepEqual(got, c.want) {
			t.Fatalf("%+v: %v %v", c.c, got, err)
		}
	}
	for _, bad := range []Claim{{Kind: ObjSpan, ObjSpan: Span{Off: 0x10000}}, {Kind: 3}} {
		if _, err := encodeClaim(nil, bad, testWire); err == nil {
			t.Fatalf("%+v: want error", bad)
		}
	}
}
