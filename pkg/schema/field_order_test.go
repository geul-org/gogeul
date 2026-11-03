//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestFieldOrder — record, raw and rest fields with fixed raw fields after the rest pass; a FIELD of no DECLARE, a second rest field and a record field after the rest fail E11; rest fields of different DECLAREs do not meet
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFieldOrder(t *testing.T) {
	f := func(id, decl uint64, w core.Word) rec {
		return recOf(core.EncodeFieldDecl(core.FieldDecl{EdgeID: id, Declare: decl, Width: w}))
	}
	decls := map[uint64]bool{1: true, 2: true}
	cases := []struct {
		body []rec
		want string
	}{
		{[]rec{f(3, 1, 0), f(4, 1, 0xFFFF), f(5, 1, 2), f(6, 2, 0xFFFF), f(7, 2, 0xFFFE), recOf(core.EncodeUintValue(1))}, ""},
		{[]rec{f(3, 9, 1)}, "E11"},
		{[]rec{f(3, 1, 0xFFFF), f(4, 1, 0xFFFF)}, "E11"},
		{[]rec{f(3, 1, 0xFFFF), f(4, 1, 1), f(5, 1, 0)}, "E11"},
	}
	for i, c := range cases {
		if got := rulesOf(fieldOrder(c.body, decls)); got != c.want {
			t.Fatalf("case %d: %q want %q", i, got, c.want)
		}
	}
}
