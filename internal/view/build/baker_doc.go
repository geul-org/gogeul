//ff:func feature=build type=command control=sequence
//ff:what baker.doc — copy a document's bytes and write its viewer page beside them, without the .geul extension
package build

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// doc copies the document rel (relative to the tree, ending in .geul) and
// writes its viewer page at the same path without .geul.
func (b *baker) doc(rel string) error {
	data, err := fs.ReadFile(b.tree, rel)
	if err != nil {
		return err
	}
	to := filepath.Join(b.out, filepath.FromSlash(rel))
	if err := os.WriteFile(to, data, 0o644); err != nil {
		return err
	}
	u := wellKnown + strings.TrimSuffix(rel, ".geul")
	err = writeFile(strings.TrimSuffix(to, ".geul"), func(w io.Writer) error { return pageHTML(w, b.tree, u, data, b.spec) })
	if err != nil {
		return fmt.Errorf("%s: %w", rel, err)
	}
	return nil
}
