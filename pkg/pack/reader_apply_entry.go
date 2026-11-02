//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.applyEntry — apply one inventory entry to a state: set (and h/ share) or delete (§7.1)
package pack

import "context"

// applyEntry applies e of the data pack obj to st.
func (r *Reader) applyEntry(ctx context.Context, st map[string]Loc, obj uint32, inv inventory, e invEntry) error {
	s, ok, err := entryStep(obj, inv, e)
	if !ok {
		return err
	}
	if s.Deleted {
		delete(st, s.Path)
		return nil
	}
	if s.Path == "" {
		p, err := r.HashPath(ctx, s.Loc)
		st[p] = s.Loc
		return err
	}
	st[s.Path] = s.Loc
	return r.applyShare(ctx, st, e, s.Loc)
}
