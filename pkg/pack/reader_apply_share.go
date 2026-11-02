//ff:func feature=pack type=resolver control=sequence
//ff:what Reader.applyShare — when flags bit0 is set, the record also serves its `h/<content hash>.geul`
package pack

import "context"

// applyShare adds the h/ path of a shared record to st.
func (r *Reader) applyShare(ctx context.Context, st map[string]Loc, e invEntry, l Loc) error {
	if e.flags&flagHash == 0 {
		return nil
	}
	p, err := r.HashPath(ctx, l)
	if err != nil {
		return err
	}
	st[p] = l
	return nil
}
