//ff:func feature=page type=helper control=selection
//ff:what tagName — the name tagBits shows for a codec: a slot marker, a 5Q Reference with its Quantifier, the spec name, fallback, or hex
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// tagName names codec of Kind k for tagBits: a slot marker by its marker name,
// a 5Q Reference as `<scheme> · <Quantifier>`, else core.CodecName, else
// fallback, else the code in hex.
func tagName(k core.Kind, codec uint16, fallback string) string {
	name := core.CodecName(k, codec)
	switch {
	case k == core.KindReference && core.IsSlotMarkerCodec(codec):
		return "자리 표지 " + core.SlotMarkerName(codec)
	case k == core.KindReference && core.QuantifierName(codec) != "":
		return name + " · " + core.QuantifierName(codec)
	case name != "":
		return name
	case fallback != "":
		return fallback
	default:
		return fmt.Sprintf("0x%03X", codec)
	}
}
