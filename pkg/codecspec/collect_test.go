//ff:func feature=codecspec type=decoder control=sequence
//ff:what TestCollect — a declared codec with named fields and enums; a DECLARE without FIELDs is empty; ISSUER targets and repeats are skipped
package codecspec

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCollect(t *testing.T) {
	body := append(sampleBody(),
		core.EncodeDeclare(core.Declare{EdgeID: 11, Target: core.TargetWord(core.KindEdge, core.EdgeTombstone)}),
		core.EncodeDeclare(core.Declare{EdgeID: 12, Target: core.TargetWord(core.KindValue, core.CodecValueIssuer2Word), Issuer: []core.Word{0, 2}, Format: 2}),
		core.EncodeDeclare(core.Declare{EdgeID: 13, Target: core.TargetWord(core.KindEdge, core.EdgeDeprecation)}),
		core.EncodeFieldDecl(core.FieldDecl{EdgeID: 14, Declare: 11, Width: 9}),
	)
	s := Collect(body)
	fs := s[core.TargetWord(core.KindEdge, core.EdgeTombstone)]
	if len(s) != 2 || len(fs) != 3 || fs[0].Name != "doc" || fs[2].Width != 1 || fs[2].Enums["0x0001"] != "license" {
		t.Fatalf("got %+v", s)
	}
	if dep, ok := s[core.TargetWord(core.KindEdge, core.EdgeDeprecation)]; !ok || len(dep) != 0 {
		t.Fatalf("DEPRECATION: %+v %v", dep, ok)
	}
}
