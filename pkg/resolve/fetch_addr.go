//ff:func feature=resolve type=resolver control=selection
//ff:what FetchAddr — the bytes of an address from a Source: stored paths as they are, latest.geul through the version list (or version 1 of a list)
package resolve

import (
	"context"
	"fmt"
)

// FetchAddr reads a from src. A version, versions.geul and a hash are stored
// paths. latest.geul is a reserved name, not stored: for a codec or document
// it is the newest live version the version list names, returned as it is
// (no redirect); for a list document (issuers/, codecs/, docs/) it is its only
// version, 1 (listDocPath). A missing address is ErrNotFound.
func FetchAddr(ctx context.Context, src Source, a Addr) ([]byte, error) {
	switch {
	case a.Kind == AddrHash || a.Versions || a.Version != Latest:
		return src.Fetch(ctx, a.Path())
	case a.Kind != AddrCodec && a.Kind != AddrDoc:
		return src.Fetch(ctx, listDocPath(a.folder()))
	}
	b, err := src.Fetch(ctx, a.folder()+"/"+listFile)
	if err != nil {
		return nil, err
	}
	k := Key{Issuer: a.Issuer, Doc: a.ID}.DocKey()
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
