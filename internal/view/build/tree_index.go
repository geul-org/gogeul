//ff:func feature=build type=helper control=sequence
//ff:what treeIndex — the listing data of the tree directory at a URL path, its entries named by treeLabels
package build

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/resolve"
)

// treeIndex lists the tree directory at URL path urlPath (ending in /) under
// root, showing Issuer and document IDs by their names.
func treeIndex(t resolve.Store, urlPath string) (indexPage, error) {
	return indexPageOf(t, treeName(urlPath), urlPath, treeLabels(t, strings.TrimSuffix(urlPath, "/")))
}
