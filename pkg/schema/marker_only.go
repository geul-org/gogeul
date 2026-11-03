//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what markerOnly — E5: a field whose CODEC list holds only slot marker targets (0000·0010·0020: always a marker, never a record)
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// markerOnly reports each field whose CODEC targets are all slot markers
// (core.IsSlotMarkerWord): such a field could never hold a record, only a
// marker (edge/CODEC 선택 필드와 자리 표지). It generalizes "0000 alone".
func markerOnly(body []rec) []Fact {
	others := map[uint64]bool{}
	last := map[uint64]rec{}
	var order []uint64
	for _, r := range body {
		c, err := core.DecodeCodecDecl(r.R)
		if err != nil {
			continue
		}
		if _, ok := last[c.Field]; !ok && !others[c.Field] {
			order = append(order, c.Field)
		}
		if !core.IsSlotMarkerWord(c.Target) {
			others[c.Field] = true
			continue
		}
		last[c.Field] = r
	}
	var fs []Fact
	for _, f := range order {
		if r, ok := last[f]; ok && !others[f] {
			fs = append(fs, factAt("E5", Error, r, "대상", "자리 표지(0000·0010·0020) 말고 대상 하나 이상", fmt.Sprintf("#%d에 자리 표지만", f)))
		}
	}
	return fs
}
