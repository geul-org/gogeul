//ff:func feature=schema type=rule control=sequence
//ff:what checkIssuerDocPointer — R5: the Issuer of ISSUER-, DOC- or CODEC-POINTER is 2/4/16 WORD and canonical; the reserved document ID F…F only under Issuer 0x00000000
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
// that holds the Issuer's value (core.IssuerCanonical). A DOC-POINTER of
// the reserved list ID F…F is Issuer 0x00000000's (its Issuer list); no other
// Issuer uses that ID (grammar 0x041L).
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
	if c == core.CodecDocPointer && core.FormatID(p[w:w+4]) == reservedListID && core.FormatID(p[:w]) != "0x00000000" {
		return []Fact{factAt("R5", Error, r, "문서 ID", "예약 문서 ID 0xFFFFFFFFFFFFFFFF는 Issuer 0x00000000의 Issuer 목록에만", core.FormatID(p[:w]))}
	}
	return nil
}
