//ff:func feature=build type=command control=selection
//ff:what baker.visit — bake one walked entry: a directory's listing or a document's bytes and page
package build

import (
	"io/fs"
	"strings"
)

// visit is the WalkDir callback: a directory gets its listing, a .geul file
// its copy and page, and other files are skipped.
func (b *baker) visit(p string, d fs.DirEntry, err error) error {
	rel := p
	switch {
	case err != nil:
		return err
	case d.IsDir():
		b.count.Dirs++
		return b.dir(strings.TrimPrefix(rel+"/", "./"))
	case strings.HasSuffix(rel, ".geul"):
		b.count.Docs++
		return b.doc(rel)
	default:
		return nil
	}
}
