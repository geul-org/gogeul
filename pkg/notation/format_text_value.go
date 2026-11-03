//ff:func feature=notation type=decoder control=selection
//ff:what formatTextValue — a byte-string Value as a token: UTF8-TEXT quoted, URI in angle brackets, others by codec name
package notation

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// formatTextValue writes a byte-string Value of even codec even holding s.
func formatTextValue(even uint16, s string) string {
	switch even {
	case core.CodecUTF8TextEven:
		return strconv.Quote(s)
	case codecURIEven:
		return "<" + s + ">"
	default:
		return core.CodecName(core.KindValue, even) + "(" + strconv.Quote(s) + ")"
	}
}
