//ff:func feature=schema type=rule control=sequence
//ff:what checkPointer — R3: EXTERNAL-POINTER IMPORT ID width, narrowest codec, its IMPORT-DOC and a target ID not 0
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkPointer applies R3 (and R2 to the target) to an EXTERNAL-POINTER-*
// Reference of scheme (C/5, 6~8). imports holds the IMPORT IDs of the HEAD.
// Target ID 0 (the imported document itself) is invalid: a document is
// pointed at only by DOC-POINTER.
func checkPointer(r rec, scheme uint16, imports map[uint64]bool) []Fact {
	w := core.ExternalPointerWidth(scheme)
	p := r.R.Payload
	if len(p) < w+1 {
		return []Fact{factAt("R3", Error, r, "Payload", fmt.Sprintf("IMPORT ID %d WORD + 대상 ID", w), fmtWords(len(p)))}
	}
	var fs []Fact
	id, err := core.DecodeID(p[:w], w)
	if err == nil && w > core.MinWidth(id) {
		fs = append(fs, factAt("R3", Error, r, "IMPORT ID", fmt.Sprintf("EXTERNAL-POINTER-%dWORD", core.MinWidth(id)), fmt.Sprintf("EXTERNAL-POINTER-%dWORD", w)))
	}
	if err != nil || !imports[id] {
		fs = append(fs, factAt("R3", Error, r, "IMPORT ID", "HEAD의 IMPORT-DOC이 배정한 IMPORT ID", fmt.Sprintf("%d", id)))
	}
	if len(p) == w+1 && p[w] == 0 {
		fs = append(fs, factAt("R3", Error, r, "대상 ID", "0이 아님 (문서 자신은 DOC-POINTER)", "0"))
	}
	return append(fs, checkTarget(r, p[w:])...)
}
