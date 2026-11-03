//ff:func feature=resolve type=resolver control=selection
//ff:what FetchAddr — the bytes of an address from a Source: stored paths as they are, latest.geul through the version list of the document, codec or list
package resolve

import (
	"context"
	"fmt"
)

// FetchAddr reads a from src. A version, versions.geul and a hash are stored
// paths. latest.geul is a reserved alias, not stored: for a document, a codec
// and a list (issuers/, codecs/, docs/ under their reserved IDs, listKey)
// alike it is the newest version the version list names that is neither
// deprecated nor deleted, returned as it is (no redirect). A missing address
// is ErrNotFound.
func FetchAddr(ctx context.Context, src Source, a Addr) ([]byte, error) {
	switch {
	case a.Kind == AddrHash || a.Versions || a.Version != Latest:
		return src.Fetch(ctx, a.Path())
	}
	k := listKey(a)
	if a.Kind == AddrCodec || a.Kind == AddrDoc {
		k = Key{Issuer: a.Issuer, Doc: a.ID}.DocKey()
	}
	b, err := src.Fetch(ctx, listPath(k))
	if err != nil {
		return nil, err
	}
	lg, err := decodeVersionList(b, k.Doc)
	if err != nil {
		return nil, err
	}
	e, ok := pickVersion(lg.Entries, Latest)
	if !ok {
		return nil, fmt.Errorf("%w: %s has no live version", ErrNotFound, a.Path())
	}
	return src.Fetch(ctx, versionPath(k, e.Version))
}
