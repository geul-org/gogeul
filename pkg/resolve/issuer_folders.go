//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what issuerFolders — the canonical Issuer folders of a local `.well-known/geul/`; none when it does not exist
package resolve

import (
	"errors"
	"io/fs"
	"os"
)

// issuerFolders lists dir's folders named as Issuer IDs.
func issuerFolders(dir string) ([]string, error) {
	ents, err := os.ReadDir(dir)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	var out []string
	for _, e := range ents {
		if e.IsDir() && isIssuerID(e.Name()) {
			out = append(out, e.Name())
		}
	}
	return out, nil
}
