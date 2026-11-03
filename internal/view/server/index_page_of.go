//ff:func feature=server type=helper control=sequence
//ff:what indexPageOf — the listing data of a tree directory: path, parent link and entries
package server

import (
	"io/fs"
	"path"
)

// indexPageOf lists the directory dir whose URL path is urlPath (ending in /).
// Up is the parent URL, empty at the top of the tree. Directories show the
// names labels gives them (Issuer ID → name, treeLabels).
func indexPageOf(t fs.FS, dir, urlPath string, labels map[string]string) (indexPage, error) {
	entries, err := indexEntries(t, dir, urlPath)
	if err != nil {
		return indexPage{}, err
	}
	labelEntries(entries, labels)
	p := indexPage{Path: urlPath, Entries: entries}
	if urlPath != wellKnown {
		p.Up = path.Dir(path.Dir(urlPath+"x")) + "/"
	}
	return p, nil
}
