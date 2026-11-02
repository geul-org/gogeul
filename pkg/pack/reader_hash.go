//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.Hash — the document of `h/<content hash>.geul`, its content hash checked against the key
package pack

import (
	"context"
	"fmt"
)

// Hash reads the record served at h/<h>.geul and verifies its GEUL content
// hash; a mismatch wraps both ErrCorrupt and ErrVerify.
func (r *Reader) Hash(ctx context.Context, h [32]byte) ([]byte, error) {
	path := hPath(string(h[:]))
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
	return b, nil
}
