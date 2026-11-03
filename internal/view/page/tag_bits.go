//ff:func feature=page type=helper control=selection
//ff:what tagBits — the bit fields of a TAG WORD 000K KKCC CCCC LLLL spelled out: prefix, KKK and Kind, C and its name (with the Quantifier of a 5Q Reference), LLLL and the length
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// tagBits spells out TAG WORD `000K KKCC CCCC LLLL`: the prefix, the 3-bit
// KKK with its Kind, the 6-bit C with the codec's name (tagName), and the
// Length code with the word count it gives (or the extension it calls for).
func tagBits(w core.Word, fallback string) string {
	k, codec := core.SplitTarget(w)
	l := uint8(w & 0xF)
	length := fmt.Sprintf("%d WORDs", l+1)
	switch l {
	case 0:
		length = "1 WORD"
	case 0xE:
		length = "Length 확장 1 WORD"
	case 0xF:
		length = "Length 확장 2 WORD"
	}
	name := tagName(k, codec, fallback)
	return fmt.Sprintf("%03b Proposal\n%03b %s\n%06b %s\n%04b %s", uint16(w)>>13, uint16(w)>>10&7, k, uint16(w)>>4&0x3F, name, l, length)
}
