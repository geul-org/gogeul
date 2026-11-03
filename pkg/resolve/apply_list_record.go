//ff:func feature=resolve type=rule control=selection
//ff:what applyListRecord — one version list BODY record: PUBLICATION, DEPRECATION or TOMBSTONE of the listed Issuer and document
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// applyListRecord applies one BODY record of the version list of document doc
// of Issuer issuer (canonical hex). A record of another document — another
// document ID or another Issuer (grammar 0x0B8L) — or any other record (a
// BIND included: a version list names no document, §1.1 102), is an error.
func applyListRecord(lg *versionLog, r core.Record, issuer, doc string) error {
	p, err := core.DecodePublication(r)
	if err != nil {
		return err
	}
	if got := core.FormatID(normalWords(p.Issuer)) + "/" + core.FormatID(p.Doc); got != issuer+"/"+doc {
		return fmt.Errorf("record for %s in the list of %s/%s", got, issuer, doc)
	}
	switch p.Codec {
	case core.EdgePublication:
		return addPublication(lg, p)
	default:
		return markVersion(lg, p)
	}
}
