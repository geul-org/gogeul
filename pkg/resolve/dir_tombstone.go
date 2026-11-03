//ff:func feature=resolve type=resolver control=sequence
//ff:what Dir.Tombstone — delete one published version: a TOMBSTONE with its reason, then its version file and h/ copy are removed
package resolve

import (
	"errors"
	"io/fs"

	"github.com/geul-org/gogeul/pkg/core"
)

// Tombstone appends a TOMBSTONE of version v with reason (사유 범주,
// LICENSE-GEUL-ORG 제5조) to the version list of document doc of issuer and
// removes the version's file and its `h/` copy: a deleted version is not
// served (edge/PUBLICATION). A version another stored version is a diff of
// is refused (basedOn) so no live version loses its base.
func (d Dir) Tombstone(issuer, doc []core.Word, v uint32, reason uint16) error {
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	if on, ok := d.basedOn(k, v); ok {
		return errors.New("geul/resolve: version " + FormatVersion(on) + " is stored as a diff on it; store that version as a snapshot first")
	}
	e, err := d.mark(issuer, doc, v, core.EdgeTombstone, reason)
	if err != nil {
		return err
	}
	for _, p := range []string{versionPath(k, v), hashPath(k.Issuer, e.Hash)} {
		if err := d.store().Remove(p); err != nil && !errors.Is(err, fs.ErrNotExist) {
			return err
		}
	}
	return nil
}
