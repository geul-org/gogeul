//ff:func feature=notation type=helper control=sequence
//ff:what tripleSpec — test fixture: TRIPLE declared as Predicate, Subject, Object (Wikidata IDs or any) and a 1-word Rank with ENUM names
package notation

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

// tripleSpec declares TRIPLE (0x0A) with a rest field at the end for tests of
// rest handling when rest is true.
func tripleSpec(rest bool) codecspec.Spec {
	pid := core.TargetWord(core.KindReference, core.CodecWikidataPID)
	fields := []codecspec.Field{
		{Name: "Predicate", Codecs: []core.Word{pid}},
		{Name: "Subject"},
		{Name: "Object"},
		{Name: "Rank", Width: 1, Enums: map[string]string{"0x0000": "normal", "0x0001": "preferred"}},
	}
	if rest {
		fields = append(fields, codecspec.Field{Name: "tail", Width: int(core.FieldRest)})
	}
	return codecspec.Spec{core.TargetWord(core.KindEdge, 0x0A): fields}
}
