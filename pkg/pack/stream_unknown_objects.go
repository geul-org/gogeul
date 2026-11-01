//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what Stream.unknownObjects — the objects of kind in o/ that the checkpoint has not closed (sealed after it, or left by an earlier run)
package pack

import (
	"os"
	"path/filepath"
	"strings"
)

// unknownObjects lists, in name order, the files in o/ with kind's extension
// that are none of the checkpoint's closed objects.
func (s *Stream) unknownObjects(kind byte) ([]string, error) {
	ents, err := os.ReadDir(filepath.Join(s.out, "o"))
	if err != nil {
		return nil, err
	}
	known := map[string]bool{}
	for _, o := range s.st.Objects {
		if o.Size > 0 {
			known[objectName(object{kind: o.Kind, hash: o.Hash})] = true
		}
	}
	ext := filepath.Ext(objectName(object{kind: kind}))
	var out []string
	for _, e := range ents {
		if strings.HasSuffix(e.Name(), ext) && !known["o/"+e.Name()] {
			out = append(out, e.Name())
		}
	}
	return out, nil
}
