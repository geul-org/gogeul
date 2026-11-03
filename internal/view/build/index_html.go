//ff:func feature=build type=renderer control=sequence
//ff:what indexHTML — the listing page of a tree directory
package build

import (
	"io"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// indexHTML writes the listing page of the tree directory at URL path urlPath
// (ending in /) in the tree t, for a static copy of the tree (build).
func indexHTML(w io.Writer, t resolve.Store, urlPath string) error {
	p, err := treeIndex(t, urlPath)
	if err != nil {
		return err
	}
	return indexTemplate.Execute(w, p)
}
