//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what TestEnumDecls — distinct codes on raw fields pass, the same code on two fields passes; an unknown field, a record field, a rest field, a wrong code width and a code declared twice on one field fail with E9
package schema

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEnumDecls(t *testing.T) {
	recs := func(rs ...core.Record) []rec {
		out := make([]rec, len(rs))
		for i, r := range rs {
			out[i] = recOf(r)
		}
		return out
	}
	field := func(id uint64, w core.Word) core.Record {
		return core.EncodeFieldDecl(core.FieldDecl{EdgeID: id, Declare: 1, Width: w})
	}
	enum := func(id, f uint64, code ...core.Word) core.Record {
		return core.EncodeEnumDecl(core.EnumDecl{EdgeID: id, Field: f, Code: code})
	}
	ok := recs(field(2, 1), field(3, 2), enum(4, 2, 0x0001), enum(5, 2, 0x0038), enum(6, 3, 0, 0x0001), enum(7, 3, 0, 0x0038))
	if fs := enumDecls(ok); len(fs) != 0 {
		t.Fatalf("ok: %+v", fs)
	}
	for i, c := range []struct {
		body []rec
		want string
	}{
		{recs(enum(4, 9, 0x0001)), "같은 문서의 FIELD"},
		{recs(field(2, 0), enum(4, 2, 0x0001)), "날값 FIELD"},
		{recs(field(2, 0xFFFF), enum(4, 2, 0x0001)), "나머지 날값에는 ENUM"},
		{recs(field(2, 1), enum(4, 2, 0, 0x0001)), "1 WORD"},
		{recs(field(2, 1), enum(4, 2, 0x0038), enum(5, 2, 0x0039), enum(6, 2, 0x0038)), "필드마다 한 번"},
	} {
		fs := enumDecls(c.body)
		if len(fs) != 1 || fs[0].Rule != "E9" || fs[0].Severity != Error || !strings.Contains(fs[0].Expected, c.want) {
			t.Fatalf("case %d: %+v", i, fs)
		}
	}
	dup := enumDecls(recs(field(2, 1), enum(4, 2, 0x0038), enum(5, 2, 0x0038)))
	if len(dup) != 1 || dup[0].Where.Record != recOf(enum(5, 2, 0x0038)).Index || dup[0].Actual != "0x0038 중복 (#2)" {
		t.Fatalf("duplicate: %+v", dup)
	}
}
