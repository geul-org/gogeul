//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.HashPath — the `h/` path a record serves, from its bytes' content hash (§4.5, §8.1)
package pack

import (
	"context"
)

// HashPath reads the record at l and names it by its content hash under its
// Issuer (hashOwner).
func (r *Reader) HashPath(ctx context.Context, l Loc) (string, error) {
	b, err := r.Record(ctx, l)
	if err != nil {
		return "", err
	}
	h, ok := contentHash(b)
	if !ok {
		return "", corrupt("hashes/ record is not a GEUL document")
	}
	return hPath(r.hashOwner(b), string(h[:])), nil
}
