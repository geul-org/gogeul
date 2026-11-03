//ff:func feature=build type=renderer control=sequence
//ff:what indexJSON — the listing of a tree directory as JSON for the tree browser
package build

import (
	"encoding/json"
	"io"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// indexJSON writes the listing of the tree directory at URL path urlPath
// (ending in /) in the tree t as the tree browser's JSON, for a static copy of
// the tree (build).
func indexJSON(w io.Writer, t resolve.Store, urlPath string) error {
	p, err := treeIndex(t, urlPath)
	if err != nil {
		return err
	}
	return json.NewEncoder(w).Encode(p)
}
