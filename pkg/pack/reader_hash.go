//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.Hash — the document of `h/<content hash>.geul`, its content hash checked against the key
package pack

import (
	"context"
	"fmt"
)

// Hash reads the record served at `<issuer>/hashes/<h>.geul` and verifies its
// GEUL content hash; a mismatch wraps both ErrCorrupt and ErrVerify. A record
// whose Issuer (hashOwner) is not issuer is not there: the hashes of one
// Issuer do not serve another's documents.
func (r *Reader) Hash(ctx context.Context, issuer string, h [32]byte) ([]byte, error) {
	path := hPath(issuer, string(h[:]))
	l, err := r.find(ctx, r.m.hLayers, string(h[:]), path)
	if err != nil {
		return nil, err
	}
	b, err := r.Record(ctx, l)
	if err != nil {
		return nil, err
	}
	if got, ok := contentHash(b); !ok || got != h {
		return nil, fmt.Errorf("%w: %s content hash mismatch (%w)", ErrCorrupt, path, ErrVerify)
	}
	if r.hashOwner(b) != issuer {
		return nil, notFound(path)
	}
	return b, nil
}
