//ff:func feature=resolve type=rule control=iteration dimension=1
//ff:what Dir.sameCodec — publishing a codec document: every served version already at that Issuer·codec ID declares the same target TAG, or it is refused (§1.1 103)
package resolve

import (
	"context"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// sameCodec refuses to publish b as a new version of codec doc (a 2-word codec
// ID; other documents pass) at k when a version already there that is not
// deleted declares another target TAG: an Issuer gives each codec ID to one
// codec, whatever its Kind (§1.1 103). b without exactly one DECLARE, and a
// version without one, are not compared.
func (d Dir) sameCodec(k DocKey, doc []core.Word, lg versionLog, b []byte) error {
	want, ok := declaredTarget(b)
	if len(doc) != 2 || !ok {
		return nil
	}
	for _, e := range lg.Entries {
		if e.Deleted {
			continue
		}
		old, err := d.Fetch(context.Background(), versionPath(k, e.Version))
		if err != nil {
			return err
		}
		if got, ok := declaredTarget(old); ok && got != want {
			return fmt.Errorf("geul/resolve: codec %s/%s version %s declares %04X, not %04X: one codec ID is one codec", k.Issuer, k.Doc, FormatVersion(e.Version), got, want)
		}
	}
	return nil
}
