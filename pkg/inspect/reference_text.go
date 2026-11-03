//ff:func feature=inspect type=decoder control=selection
//ff:what referenceText — a Reference as short text: Qn for a Wikidata item, else its codec name and payload hex
package inspect

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// referenceText names a Reference held inside a Value (a coordinate's globe,
// a quantity's unit): "Q111" for a WIKIDATA-QID that fits 64 bits, else the
// codec name and payload hex.
func referenceText(g core.Record) string {
	switch {
	case g.Codec == core.CodecWikidataQID && len(g.Payload) <= 4:
		var n uint64
		for _, w := range g.Payload {
			n = n<<16 | uint64(w)
		}
		return "Q" + strconv.FormatUint(n, 10)
	default:
		return core.CodecName(g.Kind, g.Codec) + " " + core.FormatID(g.Payload)
	}
}
