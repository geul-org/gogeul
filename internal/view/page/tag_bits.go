//ff:func feature=page type=helper control=selection
//ff:what tagBits — the bit fields of a TAG WORD 000K KKKC CCCC LLLL spelled out: prefix, section KKKK and Kind, C (offset in the section) and the codec name (with the Quantifier of a 5Q Reference), LLLL and the length
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// tagBits spells out TAG WORD `000K KKKC CCCC LLLL`: the prefix, the 4-bit
// section KKKK with its Kind (core.KindOf), the 5-bit C (offset in the
// section) with the codec's name (tagName), and the Length code with the word
// count it gives (or the extension it calls for).
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
	return fmt.Sprintf("%03b Proposal\n%04b %s\n%05b %s\n%04b %s", uint16(w)>>13, uint16(w)>>9&0xF, k, uint16(w)>>4&0x1F, name, l, length)
}
