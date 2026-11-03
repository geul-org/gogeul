//ff:func feature=resolve type=resolver control=selection
//ff:what Dir.latest — the stored bytes of the newest live version of k (latest.geul), read through its version list
package resolve

import (
	"context"
	"fmt"
)

// latest reads k's version list and returns the stored bytes of its newest
// version that is neither deprecated nor deleted (latest.geul is an alias,
// not a stored file). A document, codec or list without one is ErrNotFound.
func (d Dir) latest(k DocKey) ([]byte, error) {
	lg, err := d.readList(k, k.Doc)
	if err != nil {
		return nil, err
	}
	e, ok := pickVersion(lg.Entries, Latest)
	switch {
	case !lg.Present || !ok:
		return nil, fmt.Errorf("%w: %s has no live version", ErrNotFound, docFolder(k))
	default:
		return d.Fetch(context.Background(), versionPath(k, e.Version))
	}
}
