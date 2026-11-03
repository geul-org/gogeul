//ff:func feature=schema type=rule control=sequence
//ff:what checkSlotMarker — R1: a slot marker code (C < 3) standing as a HEAD or BODY record
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkSlotMarker applies R1 to a Reference of a slot marker code (C 0, 1, 2)
// standing as a HEAD or BODY record. A slot marker (ABSENT, UNKNOWN-VALUE,
// NO-VALUE; WORD 0x0000, 0x0010, 0x0020) is valid only inside a record field
// that names it (reference/slot-marker), so here it is an Error; a marker C
// with a payload is reserved (Warn).
func checkSlotMarker(r rec) []Fact {
	if core.IsSlotMarker(r.R) {
		name := core.SlotMarkerName(r.R.Codec)
		return []Fact{factAt("R1", Error, r, "자리 표지", "표지를 받는 레코드 필드 안", fmt.Sprintf("레코드 자리의 %s (0x%04X)", name, uint16(r.R.Codec)<<4))}
	}
	return []Fact{factAt("R1", Warn, r, "자리 표지", "Payload 없는 1 WORD", fmt.Sprintf("C %d + %d WORD (예약)", r.R.Codec, len(r.R.Payload)))}
}
