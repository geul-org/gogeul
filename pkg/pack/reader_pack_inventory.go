//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what Reader.PackInventory — hand fn one data pack's inventory step by step, entries then deleted hashes, keeping no path map
package pack

import (
	"context"
	"fmt"
)

// PackInventory calls fn for each step of data pack obj in application order;
// an error from fn stops it and is returned. Replaying DataPacks in order
// gives State. The pack's inventory is one zstd frame, read and checksummed
// whole before the first step, so it stays in memory until PackInventory
// returns; nothing else is kept and no record is read.
func (r *Reader) PackInventory(ctx context.Context, obj uint32, fn func(InventoryEntry) error) error {
	if uint64(obj) >= uint64(len(r.m.objects)) || r.m.objects[obj].kind != kindData {
		return fmt.Errorf("geul/pack: object %d is not a data pack", obj)
	}
	inv, err := r.inventory(ctx, obj)
	if err != nil {
		return err
	}
	for _, e := range inv.entries {
		s, ok, err := entryStep(obj, inv, e)
		if err != nil {
			return err
		}
		if !ok {
			continue
		}
		if err := fn(s); err != nil {
			return err
		}
	}
	for _, h := range inv.deleted {
		if err := fn(InventoryEntry{Path: hPath(h), Deleted: true}); err != nil {
			return err
		}
	}
	return nil
}
