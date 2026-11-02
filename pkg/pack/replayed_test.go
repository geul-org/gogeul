//ff:func feature=pack type=helper control=iteration dimension=1
//ff:what replayed — the state a caller rebuilds from DataPacks and PackInventory, oldest pack first
package pack

import (
	"context"
	"testing"
)

// replayed replays r's steps into a path → location map.
func replayed(t *testing.T, r *Reader) map[string]Loc {
	t.Helper()
	st := map[string]Loc{}
	for _, e := range inventorySteps(t, r) {
		if err := replayStep(context.Background(), r, st, e); err != nil {
			t.Fatal(e, err)
		}
	}
	return st
}
