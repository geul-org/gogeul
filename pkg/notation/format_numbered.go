//ff:func feature=notation type=decoder control=sequence
//ff:what formatNumbered — LOCAL as #n, WIKIDATA-QID as Qn, WIKIDATA-PID as Pn, from a minimal unsigned target ID
package notation

import (
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
)

// formatNumbered writes the target ID of Reference C c (LOCAL, QID or PID,
// Identified). An ID
// that is not minimal or is past 64 bits is false, so it stays raw.
func formatNumbered(c uint16, payload []core.Word) (string, bool) {
	n, err := core.DecodeUintValue(core.Record{Kind: core.KindValue, Codec: core.CodecUnsignedInteger, Payload: payload})
	if err != nil {
		return "", false
	}
	prefix := map[uint16]string{core.CodecLocal: "#", core.CodecWikidataQID: "Q", core.CodecWikidataPID: "P"}[c]
	return prefix + strconv.FormatUint(n, 10), true
}
