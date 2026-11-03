//ff:func feature=schema type=rule control=sequence
//ff:what sourceFacts — E8 for one SOURCE: exactly Target and Origin; Target LOCAL 0 or a BODY Edge Id; Origin a Reference, ABSENT or UNKNOWN-VALUE
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// sourceFacts checks r, a SOURCE Edge (edge/SOURCE), against ids, the Edge Ids
// of the BODY. The payload after the Edge Id is exactly two records. The
// Target is a LOCAL Reference: 0 (the whole document; such a SOURCE belongs
// in the HEAD, misplacedHeadEdges) or a BODY Edge Id. The Origin is a
// Reference; of the slot markers (Reference C < 3) it takes ABSENT (not
// written) and UNKNOWN-VALUE, not NO-VALUE.
func sourceFacts(r rec, ids map[uint64]bool) []Fact {
	target, origin, rest, ok := splitSource(r)
	if !ok || len(rest) != 0 {
		return []Fact{factAt("E8", Error, r, "Payload", "Edge Id, 대상, 원천 레코드 셋", fmtWords(len(r.R.Payload)))}
	}
	var fs []Fact
	t, _, _ := core.DecodeRecord(target)
	id, err := core.DecodeLocalRef(t)
	if err != nil || (id != 0 && !ids[id]) {
		fs = append(fs, factAt("E8", Error, r, "대상", "LOCAL 0(문서 전체) 또는 BODY의 Edge Id", fmt.Sprintf("%04X", target)))
	}
	o, _, _ := core.DecodeRecord(origin)
	marker := o.Kind == core.KindReference && core.IsSlotMarkerCodec(o.Codec)
	taken := len(origin) == 1 && (origin[0] == core.Absent || origin[0] == core.UnknownValue)
	if o.Kind != core.KindReference || (marker && !taken) {
		fs = append(fs, factAt("E8", Error, r, "원천", "Reference, ABSENT(0000) 또는 UNKNOWN-VALUE(0010)", fmt.Sprintf("%04X", origin)))
	}
	return fs
}
