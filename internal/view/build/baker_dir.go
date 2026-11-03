//ff:func feature=build type=command control=sequence
//ff:what baker.dir — write a directory's listing page and its JSON
package build

import (
	"io"
	"os"
	"path/filepath"
)

// dir writes index.html and index.json for the tree directory rel (relative to
// the tree, ending in /; empty at the top).
func (b *baker) dir(rel string) error {
	dir := filepath.Join(b.out, filepath.FromSlash(rel))
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	u := wellKnown + rel
	if err := writeFile(filepath.Join(dir, "index.html"), func(w io.Writer) error { return indexHTML(w, b.tree, u) }); err != nil {
		return err
	}
	return writeFile(filepath.Join(dir, "index.json"), func(w io.Writer) error { return indexJSON(w, b.tree, u) })
}
