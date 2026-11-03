//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestDeclareDecls — distinct targets pass, an Issuer codec apart from its ISSUER code too; a target declared twice fails E11; a FIELD of no DECLARE reaches fieldOrder
package schema

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDeclareDecls(t *testing.T) {
	decl := func(id uint64, target core.Word) rec {
		return recOf(core.EncodeDeclare(core.Declare{EdgeID: id, Target: target}))
	}
	issuerCodec := recOf(core.EncodeDeclare(core.Declare{EdgeID: 3, Target: 0x09D0, Issuer: []core.Word{0, 7}, Format: 1}))
	field := recOf(core.EncodeFieldDecl(core.FieldDecl{EdgeID: 4, Declare: 1, Width: 1}))
	cases := []struct {
		body []rec
		want string
	}{
		{[]rec{decl(1, 0x0AA0), decl(2, 0x09D0), issuerCodec, field}, ""},
		{[]rec{decl(1, 0x0AA0), decl(2, 0x0AA0)}, "E11"},
		{[]rec{decl(2, 0x0AA0), field}, "E11"},
		{[]rec{recOf(core.EncodeUintValue(1))}, ""},
	}
	for i, c := range cases {
		if got := rulesOf(declareDecls(c.body)); got != c.want {
			t.Fatalf("case %d: %q want %q", i, got, c.want)
		}
	}
}
