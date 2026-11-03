//ff:func feature=schema type=decoder control=sequence
//ff:what partKind — the Kind of one whole record cut from a payload (statement part), and whether it is a LOCAL Reference
package schema

import "github.com/geul-org/gogeul/pkg/core"

// partKind decodes the record p and returns its Kind and whether it is a
// LOCAL Reference (Quantifier Identified, core.DecodeLocalRef). p is a whole
// record splitStatement cut, so it decodes.
func partKind(p []core.Word) (core.Kind, bool) {
	r, _, _ := core.DecodeRecord(p)
	_, err := core.DecodeLocalRef(r)
	return r.Kind, err == nil
}
