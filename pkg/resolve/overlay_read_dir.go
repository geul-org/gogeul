//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what overlay.ReadDir — fs.ReadDirFS: the union of the stored and loose entries in name order, without the pack sets' own files
package resolve

import (
	"errors"
	"io/fs"
	"slices"
	"strings"
)

// ReadDir lists name. It is missing only when both sides lack it; the stored
// entry wins a name both have.
func (o overlay) ReadDir(name string) ([]fs.DirEntry, error) {
	out, err := o.stored.ReadDir(name)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	loose, lerr := o.loose.ReadDir(name)
	if lerr != nil && (err != nil || !errors.Is(lerr, fs.ErrNotExist)) {
		return nil, lerr
	}
	seen := map[string]bool{setLatest: true, setManifests: true, setObjects: true}
	for _, e := range out {
		seen[e.Name()] = true
	}
	for _, e := range loose {
		if !seen[e.Name()] && !strings.HasPrefix(e.Name(), ".") {
			out = append(out, e)
		}
	}
	slices.SortFunc(out, func(a, b fs.DirEntry) int { return strings.Compare(a.Name(), b.Name()) })
	return out, nil
}
