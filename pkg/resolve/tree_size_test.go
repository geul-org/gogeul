//ff:func feature=resolve type=helper control=sequence
//ff:what treeSize — test fixture: the bytes of a data folder's files
package resolve

import (
	"io/fs"
	"path/filepath"
	"testing"
)

// treeSize sums the files under root.
func treeSize(t *testing.T, root string) int64 {
	t.Helper()
	var n int64
	err := filepath.WalkDir(root, func(_ string, e fs.DirEntry, err error) error {
		if err != nil || e.IsDir() {
			return err
		}
		fi, err := e.Info()
		n += fi.Size()
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return n
}
