//ff:func feature=pack type=helper control=sequence
//ff:what replayStep — apply one InventoryEntry to a state through the public API alone
package pack

import "context"

// replayStep does what the InventoryEntry documentation says a caller does.
func replayStep(ctx context.Context, r *Reader, st map[string]Loc, e InventoryEntry) error {
	if e.Deleted {
		delete(st, e.Path)
		return nil
	}
	if e.Path != "" {
		st[e.Path] = e.Loc
	}
	if !e.Hashed {
		return nil
	}
	p, err := r.HashPath(ctx, e.Loc)
	st[p] = e.Loc
	return err
}
