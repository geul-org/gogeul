//ff:func feature=build type=helper control=iteration dimension=1
//ff:what Write — copy the embedded third-party scripts into a directory, for a static copy of the tree
package static

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Write copies the embedded vendor files into dir, creating it, so a static
// copy of the tree serves them at `_static/`, where the page loads them (build).
func Write(dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	des, _ := fs.ReadDir(files, "vendor")
	for _, de := range des {
		b, _ := files.ReadFile("vendor/" + de.Name())
		if err := os.WriteFile(filepath.Join(dir, de.Name()), b, 0o644); err != nil {
			return err
		}
	}
	return nil
}
