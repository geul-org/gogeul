//ff:func feature=page type=helper control=selection
//ff:what refText — a LOCAL or Wikidata Reference (Quantifier Identified) as #n, Qn or Pn
package page

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// refText shows a Reference that names one thing by number: LOCAL as #n,
// WIKIDATA-QID as Qn and WIKIDATA-PID as Pn, n read from the payload words
// big-endian. Other Reference codecs, other Quantifiers, an empty or over 4
// word payload and non-Reference records are false.
func refText(r core.Record) (string, bool) {
	if r.Kind != core.KindReference || len(r.Payload) == 0 || len(r.Payload) > 4 {
		return "", false
	}
	n := wordsNumber(r.Payload)
	switch r.Codec {
	case core.CodecLocal:
		return "#" + strconv.FormatUint(n, 10), true
	case core.CodecWikidataQID:
		return "Q" + strconv.FormatUint(n, 10), true
	case core.CodecWikidataPID:
		return "P" + strconv.FormatUint(n, 10), true
	default:
		return "", false
	}
}
