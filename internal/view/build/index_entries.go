//ff:func feature=build type=helper control=iteration dimension=1
//ff:what indexEntries — the subdirectories and .geul files of a directory as links
package build

import (
	"io/fs"
	"strings"
)

// indexEntries lists the tree directory dir's subdirectories (with a trailing /) and .geul files,
// in name order, as links under the URL path base (ending in /). A document
// links to its path without .geul, where build puts its viewer page.
func indexEntries(t fs.FS, dir, base string) ([]indexEntry, error) {
	des, err := fs.ReadDir(t, dir)
	if err != nil {
		return nil, err
	}
	var out []indexEntry
	for _, de := range des {
		name := de.Name()
		if de.IsDir() {
			name += "/"
		}
		if de.IsDir() || strings.HasSuffix(name, ".geul") {
			out = append(out, indexEntry{Name: name, URL: base + strings.TrimSuffix(name, ".geul")})
		}
	}
	return out, nil
}
