//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what enumDecls — E9: each ENUM names a fixed raw FIELD of the document (not a rest field), its code is exactly that field's width, and no field declares the same code twice
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// enumDecls checks the ENUM Edges of body against its FIELDs (edge/ENUM): the
// field is a FIELD Edge of this document with width n of 1~65534 (a record
// field and a rest field 0xFFFF take no ENUM, grammar 0x0BCL), the code is
// exactly n words, and one field never declares the same code twice. ENUMs that do not decode are E4's.
func enumDecls(body []rec) []Fact {
	widths := map[uint64]core.Word{}
	for _, r := range body {
		if f, err := core.DecodeFieldDecl(r.R); err == nil {
			widths[f.EdgeID] = f.Width
		}
	}
	seen := map[string]bool{}
	var fs []Fact
	for _, r := range body {
		e, err := core.DecodeEnumDecl(r.R)
		if err != nil {
			continue
		}
		w, ok := widths[e.Field]
		key := fmt.Sprintf("%d %s", e.Field, core.FormatID(e.Code))
		switch {
		case !ok:
			fs = append(fs, factAt("E9", Error, r, "필드", "같은 문서의 FIELD", fmt.Sprintf("#%d", e.Field)))
		case w == 0:
			fs = append(fs, factAt("E9", Error, r, "필드", "날값 FIELD(폭 ≥ 1)", fmt.Sprintf("#%d 레코드 필드", e.Field)))
		case w == 0xFFFF:
			fs = append(fs, factAt("E9", Error, r, "필드", "고정 폭 날값 FIELD(나머지 날값에는 ENUM을 붙이지 않는다)", fmt.Sprintf("#%d 나머지 날값", e.Field)))
		case len(e.Code) != int(w):
			fs = append(fs, factAt("E9", Error, r, "코드값", fmt.Sprintf("%d WORD", w), fmt.Sprintf("%d WORD", len(e.Code))))
		case seen[key]:
			fs = append(fs, factAt("E9", Error, r, "코드값", "필드마다 한 번", fmt.Sprintf("%s 중복 (#%d)", core.FormatID(e.Code), e.Field)))
		}
		seen[key] = true
	}
	return fs
}
