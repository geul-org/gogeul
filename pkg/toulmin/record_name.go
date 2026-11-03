//ff:func feature=toulmin type=helper control=sequence
//ff:what RecordName — the spec name of a TOULMIN direct codec (core.ToulminName)
package toulmin

import "github.com/geul-org/gogeul/pkg/core"

// RecordName returns the spec name of ref (grammar README 코드 색인, sections
// 8 TOULMIN and 9 TOULMIN 형식), or "" when ref is no TOULMIN record type.
// The names live in core (core.ToulminName) so core.CodeName names every
// assigned code.
func RecordName(ref CodecRef) string {
	return core.ToulminName(ref.Codec)
}
