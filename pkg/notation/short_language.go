//ff:func feature=notation type=encoder control=selection
//ff:what shortLanguage — a LANGUAGE-SHORT-CODE record from a registry tag (case ignored) or a 1-WORD hex code (assigned, when the registry is at hand)
package notation

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// shortLanguage reads the language after `@@`: `0xNNNN` is the code itself
// (1 WORD, not 0x0000 or 0xffff), any other text a BCP 47 tag the GEUL
// Language Registry in ctx.Spec (codecspec.Spec.Languages) numbers, compared
// without case. A tag the registry does not have is an error: it is written
// as a LANGUAGE-BCP47-CODE (`@tag`) instead. The registry is a closed list, so
// with a registry at hand a hex code it does not assign is an error too.
func shortLanguage(s string, ctx Context) (core.Record, error) {
	langs := ctx.Spec.Languages()
	if strings.HasPrefix(s, "0x") {
		ws, err := parseHex(s)
		_, assigned := langs[core.FormatID(ws)]
		if err != nil || len(ws) != 1 || ws[0] == 0 || ws[0] == 0xFFFF || langs != nil && !assigned {
			return core.Record{}, fmt.Errorf("geul/notation: %q is not an assigned LANGUAGE-SHORT-CODE (1 WORD in the GEUL Language Registry, never 0x0000 or 0xFFFF)", s)
		}
		return core.Record{Kind: core.KindValue, Codec: codecLanguageShort, Payload: ws}, nil
	}
	ws, err := enumWords(s, langs)
	switch {
	case s == "":
		return core.Record{}, fmt.Errorf("geul/notation: no language after @@")
	case err != nil:
		return core.Record{}, fmt.Errorf("geul/notation: %q has no number in the GEUL Language Registry; write it as @%s (LANGUAGE-BCP47-CODE)", s, s)
	default:
		return core.Record{Kind: core.KindValue, Codec: codecLanguageShort, Payload: ws}, nil
	}
}
