//ff:func feature=core type=encoder control=iteration dimension=1
//ff:what FormatID — a multi-word id as one `0x` number whose digit count encodes its width
package core

import (
	"fmt"
	"strings"
)

// FormatID renders ws as a single `0x` hex number with four digits per word, so
// the digit count is the width (GEUL-ORG ID 표기): 1 WORD `0x0001` and 2 WORD
// `0x00000001` stay distinct.
func FormatID(ws []Word) string {
	var b strings.Builder
	b.WriteString("0x")
	for _, w := range ws {
		fmt.Fprintf(&b, "%04X", uint16(w))
	}
	return b.String()
}
