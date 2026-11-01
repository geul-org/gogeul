//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what removeTemps — delete the temp files writeFileSync leaves in a folder when it is killed
package pack

import (
	"os"
	"path/filepath"
	"strings"
)

// removeTemps removes every .tmp-* file directly in dir. A process killed
// inside writeFileSync leaves one; nothing refers to it.
func removeTemps(dir string) error {
	ents, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, e := range ents {
		if !strings.HasPrefix(e.Name(), ".tmp-") {
			continue
		}
		if err := os.Remove(filepath.Join(dir, e.Name())); err != nil {
			return err
		}
	}
	return nil
}
