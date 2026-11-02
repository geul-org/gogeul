//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.State — the archive state rebuilt from data pack inventories alone, oldest pack first (§7.1)
package pack

import "context"

// State replays every data pack's inventory in OBJECTS order: records set
// their path (and h/ path when shared), deletions remove paths. The result
// must equal what the layers answer (V3).
func (r *Reader) State(ctx context.Context) (map[string]Loc, error) {
	st := map[string]Loc{}
	for _, obj := range r.DataPacks() {
		if err := r.applyPack(ctx, st, obj); err != nil {
			return nil, err
		}
	}
	return st, nil
}
