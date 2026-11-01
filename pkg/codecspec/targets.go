//ff:func feature=codecspec type=decoder control=iteration dimension=1
//ff:what Targets — the target words whose spec documents lay out these records, in order: each direct codec, and the language registry's when a language code may appear
package codecspec

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// languageTag is the TAG WORD of a one-word LANGUAGE-SHORT-CODE record. It is
// embedded in a UTF8-LANGUAGE-TEXT, which may itself be embedded, so it is
// looked for among the payload words; a data word that happens to equal it
// only reads one spec document more.
const languageTag core.Word = 0x08C8

// Targets lists what decoding records needs from the codec specs: the target
// word of each record (Spec.Of; ISSUER records are not looked up, and the
// reserved Reference codec 0000 has no spec), and the LANGUAGE-SHORT-CODE
// codec when a payload may carry a language code, since Spec.Languages names
// it though that record is not a top-level record. Sorted, without repeats.
func Targets(records []core.Record) []core.Word {
	var out []core.Word
	for _, r := range records {
		if t := keyOf(r); len(r.Issuer) == 0 && t != 0 {
			out = append(out, t)
		}
		if slices.Contains(r.Payload, languageTag) {
			out = append(out, languageCodec)
		}
	}
	slices.Sort(out)
	return slices.Compact(out)
}
