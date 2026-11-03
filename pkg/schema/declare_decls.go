//ff:func feature=schema type=rule control=iteration dimension=1
//ff:what declareDecls — E11: one DECLARE per target in a document, and each FIELD's DECLARE, rest field and the fields after it (fieldOrder)
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// declareDecls checks the DECLARE Edges of body (grammar 0x0BBL "한 문서에서
// 같은 대상을 두 번 선언하지 않는다"): a target — the TAG word, and for an Issuer
// codec its Issuer and codec number — is declared once. The FIELDs are then
// checked against those DECLAREs (fieldOrder). DECLAREs that do not decode
// are E4's.
func declareDecls(body []rec) []Fact {
	decls := map[uint64]bool{}
	targets := map[string]bool{}
	var fs []Fact
	for _, r := range body {
		d, err := core.DecodeDeclare(r.R)
		key := fmt.Sprintf("%04X %s %08X", d.Target, core.FormatID(d.Issuer), d.Format)
		if err == nil && targets[key] {
			fs = append(fs, factAt("E11", Error, r, "대상", "대상마다 DECLARE 하나", fmt.Sprintf("%04X 중복", d.Target)))
		}
		if err == nil {
			decls[d.EdgeID], targets[key] = true, true
		}
	}
	return append(fs, fieldOrder(body, decls)...)
}
