//ff:func feature=codecspec type=helper control=sequence
//ff:what keyOf — the Spec key of a record: its target word, Kind and Codec as they are (a 5Q code keeps its Quantifier)
package codecspec

import "github.com/geul-org/gogeul/pkg/core"

// keyOf is the Spec key of r: its Kind and Codec as a target word (LLLL 0).
// Every code has its own codec document, a 5Q code included: LOCAL
// Identified is 0x0050 and LOCAL Generic 0x0060 (grammar codec-spec.md
// "Reference 5Q 체계도 양화 코드마다 따로 둔다", 0x0BBL DECLARE 유효성), so a
// quantified Reference is not folded onto its scheme's Identified code.
func keyOf(r core.Record) core.Word {
	return core.TargetWord(r.Kind, r.Codec)
}
