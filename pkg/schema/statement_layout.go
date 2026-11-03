//ff:func feature=schema type=rule control=selection
//ff:what statementLayout — E10: a TRIPLE or QUALIFIER has three whole records after its Edge Id, Reference Predicate·Subject (TRIPLE) or LOCAL Target (QUALIFIER), a Reference or Value last record, then exactly one Rank WORD 0·1·2 (TRIPLE) or nothing (QUALIFIER)
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// statementLayout applies E10 to a direct-code TRIPLE or QUALIFIER (grammar
// 0x0AAL, 0x0ABL "해석과 유효성"): after the Edge Id come three whole
// records; TRIPLE's Predicate and Subject are References and after the
// Object exactly one Rank WORD of 0 (normal), 1 (preferred) or 2
// (deprecated) remains; QUALIFIER's Target is a LOCAL Reference and the
// Value ends the payload. The Object (Value) is a Reference or a Value
// record. Slot markers in these fields are E7's; a missing Edge Id is E1's.
// QUALIFIER's Property kind is not checked: the grammar's text says Reference
// but its field table takes any record but a slot marker.
func statementLayout(r rec) []Fact {
	if r.R.Kind != core.KindEdge || (r.R.Codec != core.EdgeTriple && r.R.Codec != core.EdgeQualifier) {
		return nil
	}
	_, hasID := edgeID(r)
	s, ok := splitStatement(r)
	triple := r.R.Codec == core.EdgeTriple
	k0, _ := partKind(s.Parts[0])
	k1, local := partKind(s.Parts[1])
	k2, _ := partKind(s.Parts[2])
	switch {
	case !hasID:
		return nil
	case !ok:
		return []Fact{factAt("E10", Error, r, "Payload", "Edge Id 뒤 완전한 레코드 셋", "레코드로 나뉘지 않음")}
	case triple && (k0 != core.KindReference || k1 != core.KindReference):
		return []Fact{factAt("E10", Error, r, "Predicate·Subject", "Reference 레코드", fmt.Sprintf("Kind %d, %d", k0, k1))}
	case !triple && !local:
		return []Fact{factAt("E10", Error, r, "Target", "LOCAL Reference(같은 문서의 Edge Id)", fmt.Sprintf("%04X", s.Parts[1]))}
	case k2 != core.KindReference && k2 != core.KindValue:
		return []Fact{factAt("E10", Error, r, "Object", "Reference·Value 레코드 또는 자리 표지", fmt.Sprintf("Kind %d", k2))}
	case triple && len(s.Rest) != 1:
		return []Fact{factAt("E10", Error, r, "Rank", "Object 뒤 정확히 1 WORD", fmt.Sprintf("%d WORD", len(s.Rest)))}
	case triple && s.Rest[0] > 2:
		return []Fact{factAt("E10", Error, r, "Rank", "0(normal)·1(preferred)·2(deprecated)", fmt.Sprintf("%d", s.Rest[0]))}
	case !triple && len(s.Rest) != 0:
		return []Fact{factAt("E10", Error, r, "Payload", "Value에서 끝남", fmt.Sprintf("%d WORD 남음", len(s.Rest)))}
	default:
		return nil
	}
}
