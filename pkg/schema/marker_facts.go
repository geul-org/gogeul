//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what markerFacts — E7 for one statement: no Reference Codec 0000 in the first two fields, and in the last only UNKNOWN-VALUE or NO-VALUE
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// markerFacts reports each field of s holding a Reference of a slot marker C (C < 3) that
// the field does not take: Predicate·Subject (Property·Target) take none, the
// Object (Value) takes exactly 0x0010 and 0x0020 — not ABSENT (the field is
// not optional), not a reserved marker, not a marker C with payload.
func markerFacts(s statement) []Fact {
	names := [3]string{"Predicate", "Subject", "Object"}
	if s.Codec == core.EdgeQualifier {
		names = [3]string{"Property", "Target", "Value"}
	}
	var fs []Fact
	for i, p := range s.Parts {
		r, _, _ := core.DecodeRecord(p)
		named := i == 2 && len(p) == 1 && (p[0] == core.UnknownValue || p[0] == core.NoValue)
		if r.Kind != core.KindReference || !core.IsSlotMarkerCodec(r.Codec) || named {
			continue
		}
		want := "자리 표지 없음"
		if i == 2 {
			want = "UNKNOWN-VALUE(0010) 또는 NO-VALUE(0020)만"
		}
		fs = append(fs, factAt("E7", Error, s.R, names[i], want, fmt.Sprintf("%04X", p)))
	}
	return fs
}
