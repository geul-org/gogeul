//ff:func feature=resolve type=rule control=selection
//ff:what applyListRecord — one version list BODY record: PUBLICATION, DEPRECATION or TOMBSTONE
package resolve

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// applyListRecord applies one BODY record of the version list of document doc
// (canonical hex). A record of another document, or any other record (a BIND
// included: a version list names no document, §1.1 102), is an error.
func applyListRecord(lg *versionLog, r core.Record, doc string) error {
	p, err := core.DecodePublication(r)
	if err != nil {
		return err
	}
	if core.FormatID(p.Doc) != doc {
		return fmt.Errorf("record for %s in the list of %s", core.FormatID(p.Doc), doc)
	}
	switch p.Codec {
	case core.EdgePublication:
		return addPublication(lg, p)
	default:
		return markVersion(lg, p)
	}
}
