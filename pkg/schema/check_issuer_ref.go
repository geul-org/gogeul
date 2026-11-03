//ff:func feature=schema type=rule control=sequence
//ff:what checkIssuerRef — R4: EXTERNAL-ISSUER Issuer width and canonical form and a target ID not 0
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkIssuerRef applies R4 (and R2 to the target) to an EXTERNAL-ISSUER-*
// Reference of scheme (C/5, 9~11), whose Issuer must be in canonical form
// (core.IssuerCanonical). It points at a LOCAL ID of an Issuer's
// document (document ID 0 is the Issuer's term definition index document):
// the document itself is DOC-POINTER, so target ID 0 is invalid.
func checkIssuerRef(r rec, scheme uint16) []Fact {
	w := core.ExternalIssuerWidth(scheme)
	p := r.R.Payload
	if len(p) < w+5 {
		return []Fact{factAt("R4", Error, r, "Payload", fmt.Sprintf("Issuer %d WORD + 문서 ID 4 WORD + 대상 ID", w), fmtWords(len(p)))}
	}
	fs := checkTarget(r, p[w+4:])
	if !core.IssuerCanonical(p[:w]) {
		fs = append(fs, factAt("R4", Error, r, "Issuer", "값이 들어가는 가장 짧은 폭", fmt.Sprintf("EXTERNAL-ISSUER-%dWORD", w)))
	}
	if len(p) == w+5 && p[w+4] == 0 {
		fs = append(fs, factAt("R4", Error, r, "대상 ID", "0이 아님 (문서 자신은 DOC-POINTER)", "0"))
	}
	return fs
}
