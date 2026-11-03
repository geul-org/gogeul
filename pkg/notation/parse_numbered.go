//ff:func feature=notation type=encoder control=sequence
//ff:what parseNumbered — #n, #label, Qn or Pn as a LOCAL or Wikidata Reference (LOCAL through the label map)
package notation

import (
	"fmt"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// parseNumbered reads the prefix (#, Q or P) and a decimal target ID. After
// #, a label ctx.Local knows gives its ID instead.
func parseNumbered(tok string, ctx Context) (core.Record, error) {
	c := map[byte]uint16{'#': core.CodecLocal, 'Q': core.CodecWikidataQID, 'P': core.CodecWikidataPID}[tok[0]]
	n, err := strconv.ParseUint(tok[1:], 10, 64)
	if c == core.CodecLocal && ctx.Local != nil {
		if id, ok := ctx.Local(tok[1:]); ok {
			n, err = id, nil
		}
	}
	if err != nil {
		return core.Record{}, fmt.Errorf("geul/notation: %q needs a decimal number after %c", tok, tok[0])
	}
	return core.Record{Kind: core.KindReference, Codec: c, Payload: core.EncodeUintValue(n).Payload}, nil
}
