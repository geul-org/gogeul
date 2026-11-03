//ff:func feature=notation type=encoder control=iteration dimension=1
//ff:what codecOf — the direct codec of a Kind whose core spec name is name (case ignored); Reference names give CCCC with Quantifier 0
package notation

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// codecOf looks name up among the core codec names of kind (core.CodecName).
// ISSUER codecs, which need an Issuer, are not direct and are skipped. A
// Reference scheme name gives the scheme's Identified head (the first C of
// the name), so a Quantifier other than Identified stays raw.
func codecOf(kind core.Kind, name string) (uint16, bool) {
	if name == "" {
		return 0, false
	}
	up := strings.ToUpper(name)
	for codec := uint16(0); codec <= core.CodecMax(kind); codec++ {
		if core.IssuerWidth(kind, codec) == 0 && core.CodecName(kind, codec) == up {
			return codec, true
		}
	}
	return 0, false
}
