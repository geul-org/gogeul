//ff:func feature=resolve type=rule control=iteration dimension=1
//ff:what checkDeclareKind — a codec document found for a record declares a codec of that record's Kind; another Kind is an error (§1.1 103)
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// checkDeclareKind holds every DECLARE in the codec document doc, found for a
// record of Kind want, to a target of that Kind. An Issuer assigns its codec
// IDs across Kinds in one `codecs/` (§1.1 103), so the address alone does not
// say the Kind; a document declaring another Kind is not this record's codec.
func checkDeclareKind(doc inspect.Document, want core.Kind) error {
	for _, e := range doc.Body {
		dc, err := core.DecodeDeclare(e.Record)
		if err != nil {
			continue
		}
		if got := core.KindOf(dc.Target); got != want {
			return fmt.Errorf("geul/resolve: codec document declares %04X (Kind %d) for a record of Kind %d", dc.Target, got, want)
		}
	}
	return nil
}
