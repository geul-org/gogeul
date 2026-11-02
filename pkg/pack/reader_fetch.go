//ff:func feature=pack type=resolver control=selection
//ff:what Reader.Fetch — resolve.Source: an archive path's bytes through P, H or the named table
package pack

import "context"

// Fetch returns the bytes of path (relative to `.well-known/geul/`), or an
// error wrapping ErrNotFound when this generation does not have it.
func (r *Reader) Fetch(ctx context.Context, path string) ([]byte, error) {
	ref := parsePath(path)
	switch ref.space {
	case spaceP:
		return r.Lookup(ctx, ref.issuer, ref.doc, ref.slot)
	case spaceH:
		var h [32]byte
		copy(h[:], ref.hash)
		return r.Hash(ctx, h)
	default:
		return r.named(ctx, path)
	}
}
