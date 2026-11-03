//ff:func feature=resolve type=resolver control=sequence
//ff:what Fetch — the pure GEUL bytes at an address (geul get): ParseAddr's place, then FetchAddr
package resolve

import "context"

// Fetch returns the bytes at a (geul get = ParseAddr → Fetch → standard
// output), read from the place a.Base names. A missing document is ErrNotFound.
func Fetch(ctx context.Context, a Addr) ([]byte, error) {
	src, err := sourceOfBase(a.Base)
	if err != nil {
		return nil, err
	}
	return FetchAddr(ctx, src, a)
}
