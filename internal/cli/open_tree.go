//ff:func feature=cli type=helper control=sequence
//ff:what openTree — the archive tree under a root (directory, pack sets, or pack sets with loose files beside them) for serve and build
package cli

import (
	"context"
	"path/filepath"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// openTree opens root/.well-known/geul as a Store and returns where it is stored.
func openTree(root string) (resolve.Store, string, error) {
	wk := filepath.Join(root, ".well-known", "geul")
	if !pack.IsPacked(wk) {
		return resolve.Dir{Root: root}.Tree(), wk, nil
	}
	a, err := pack.OpenDir(context.Background(), wk, packOptions())
	if err != nil {
		return nil, wk, err
	}
	return resolve.Dir{Root: root, Store: a}.Tree(), wk, nil
}
