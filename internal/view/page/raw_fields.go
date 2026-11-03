//ff:func feature=page type=helper control=sequence
//ff:what rawFields — the Issuer and format number of an uninterpreted ISSUER record
package page

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// rawFields shows what an uninterpreted record still says: the Issuer and
// 2-word codec number of an ISSUER codec. Other records show only their words.
func rawFields(r core.Record) []Field {
	if len(r.Issuer) == 0 {
		return nil
	}
	return []Field{{Key: "issuer", Text: core.FormatID(r.Issuer)}, {Key: "format", Text: fmt.Sprintf("0x%08X", r.Format)}}
}
