//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what fieldOrder — E11: a FIELD names a DECLARE of the document, a DECLARE has at most one rest field (0xFFFF), and only fixed raw fields follow it
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// fieldOrder checks each FIELD of body in document order (grammar 0x0BCL,
// codec-spec.md): its DECLARE is a DECLARE Edge of this document (decls); a
// DECLARE takes at most one rest field (width 0xFFFF), and every field after
// it is a fixed raw field (width 1~65534), because a record or another rest
// field after it could not find its end. FIELDs that do not decode are E4's.
func fieldOrder(body []rec, decls map[uint64]bool) []Fact {
	rest := map[uint64]bool{}
	var fs []Fact
	for _, r := range body {
		f, err := core.DecodeFieldDecl(r.R)
		switch {
		case err != nil:
		case !decls[f.Declare]:
			fs = append(fs, factAt("E11", Error, r, "DECLARE", "같은 문서의 DECLARE", fmt.Sprintf("#%d", f.Declare)))
		case f.Width == 0xFFFF && rest[f.Declare]:
			fs = append(fs, factAt("E11", Error, r, "폭", "DECLARE마다 나머지 날값(0xFFFF) 하나", "둘째 나머지 날값"))
		case f.Width == 0 && rest[f.Declare]:
			fs = append(fs, factAt("E11", Error, r, "폭", "나머지 날값 뒤에는 고정 폭 날값만", "레코드 필드"))
		}
		rest[f.Declare] = rest[f.Declare] || (err == nil && f.Width == 0xFFFF)
	}
	return fs
}
