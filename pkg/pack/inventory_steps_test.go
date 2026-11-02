//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what inventorySteps — every step of every data pack, in replay order
package pack

import (
	"context"
	"testing"
)

// inventorySteps walks DataPacks and PackInventory.
func inventorySteps(t *testing.T, r *Reader) []InventoryEntry {
	t.Helper()
	var out []InventoryEntry
	for _, obj := range r.DataPacks() {
		err := r.PackInventory(context.Background(), obj, func(e InventoryEntry) error {
			out = append(out, e)
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	return out
}
