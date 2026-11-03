//ff:func feature=schema type=rule control=sequence
//ff:what checkIssuerDocPointer — R5: the Issuer of ISSUER-, DOC- or CODEC-POINTER is 2/4/16 WORD and canonical
package schema

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkIssuerDocPointer applies R5 to an ISSUER-POINTER, DOC-POINTER or
// CODEC-POINTER Reference of C c. The Issuer width is what the Length leaves
// after the fixed fields: nothing for ISSUER-POINTER, the 4-word document ID
// and 2-word version ID for DOC-POINTER, the 2-word codec ID and 2-word version
// ID for CODEC-POINTER. It must be 2, 4 or 16 words and the shortest width
// that holds the Issuer's value (core.IssuerCanonical).
func checkIssuerDocPointer(r rec, c uint16) []Fact {
	p := r.R.Payload
	fixed := map[uint16]int{core.CodecDocPointer: 6, core.CodecCodecPointer: 4}[c]
	w := len(p) - fixed
	if w != 2 && w != 4 && w != 16 {
		return []Fact{factAt("R5", Error, r, "Issuer", "2/4/16 WORD", fmtWords(w))}
	}
	if !core.IssuerCanonical(p[:w]) {
		return []Fact{factAt("R5", Error, r, "Issuer", "값이 들어가는 가장 짧은 폭", fmt.Sprintf("%d WORD", w))}
	}
	return nil
}
