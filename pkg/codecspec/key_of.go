//ff:func feature=codecspec type=helper control=sequence
//ff:what keyOf — the Spec key of a record: its target word, a quantified Reference folded onto its scheme head
package codecspec

import "github.com/geul-org/gogeul/pkg/core"

// keyOf is the Spec key of r: its Kind and Codec, a quantified Reference (5 <=
// C < 0x3C) folded onto the Identified head of its scheme (edge/DECLARE).
// Slot markers, pointers, CODEC-POINTER and REFERENCE-ISSUER keep their C.
func keyOf(r core.Record) core.Word {
	c := r.Codec
	if r.Kind == core.KindReference {
		c = core.RefIdentified(c)
	}
	return core.TargetWord(r.Kind, c)
}
