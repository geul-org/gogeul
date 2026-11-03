//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.applyPack — apply one data pack's inventory entries, then its deleted h/ hashes
package pack

import "context"

// applyPack applies the inventory of data pack obj to st.
func (r *Reader) applyPack(ctx context.Context, st map[string]Loc, obj uint32) error {
	inv, err := r.inventory(ctx, obj)
	if err != nil {
		return err
	}
	for _, e := range inv.entries {
		if err := r.applyEntry(ctx, st, obj, inv, e); err != nil {
			return err
		}
	}
	for _, h := range inv.deleted {
		delete(st, hPath(r.hashOwner(nil), h))
	}
	return nil
}
