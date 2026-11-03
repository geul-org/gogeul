//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what Dir.removeOtherVersions — deletes a document's version files other than version 1
package resolve

import (
	"errors"
	"io/fs"
	"path"
	"strings"
)

// removeOtherVersions deletes every `0x????????.geul` version file in the
// folder of document k except version 1 (Overwrite). The version list and
// other files stay; a document with no folder has nothing to delete.
func (d Dir) removeOtherVersions(k DocKey) error {
	dir := docFolder(k)
	ents, err := d.store().ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	keep := FormatVersion(1) + ".geul"
	for _, e := range ents {
		if ok, _ := path.Match("0x????????.geul", e.Name()); !ok || strings.EqualFold(e.Name(), keep) {
			continue
		}
		if err := d.store().Remove(dir + "/" + e.Name()); err != nil {
			return err
		}
	}
	return nil
}
