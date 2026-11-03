//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.mark — append a DEPRECATION or TOMBSTONE for one published version to its version list
package resolve

import (
	"fmt"
	"time"

	"github.com/geul-org/gogeul/pkg/core"
)

// mark appends a DEPRECATION or TOMBSTONE (codec; reason for TOMBSTONE) of
// version v of document doc to its version list, as a new list whose PARENT is
// the old one (edge/PUBLICATION). The version must be published and not
// deleted; deprecating twice is refused. It returns the old entry.
func (d Dir) mark(issuer, doc []core.Word, v uint32, codec uint16, reason uint16) (VersionEntry, error) {
	k := Key{Issuer: issuer, Doc: doc}.DocKey()
	lg, err := d.readList(k, k.Doc)
	if err != nil {
		return VersionEntry{}, err
	}
	e, ok := entryOf(lg.Entries, v)
	switch {
	case !ok:
		return VersionEntry{}, fmt.Errorf("geul/resolve: version %s of %s/%s is not published", FormatVersion(v), k.Issuer, k.Doc)
	case e.Deleted:
		return VersionEntry{}, fmt.Errorf("geul/resolve: version %s is deleted", FormatVersion(v))
	case codec == core.EdgeDeprecation && e.Deprecated:
		return VersionEntry{}, fmt.Errorf("geul/resolve: version %s is already deprecated", FormatVersion(v))
	}
	rec := core.EncodePublication(core.Publication{Codec: codec, EdgeID: uint64(len(lg.Records)) + 1, Issuer: issuer, Doc: doc, Version: v, Reason: reason})
	list := encodeVersionList(append(lg.Records, rec), &lg.Hash, time.Now().Unix(), lg.Next)
	return e, d.writeAll(map[string][]byte{hashPath(k.Issuer, listHash(list)): list, listPath(k): list})
}
