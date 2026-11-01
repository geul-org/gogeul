//ff:func feature=cli type=helper control=sequence
//ff:what newHarness — test fixture: an empty data folder (found from cwd) and an empty config folder
package cli

import (
	"os"
	"path/filepath"
	"testing"
)

// newHarness makes the folders; the data folder holds `.well-known/geul/`
// as a directory archive (GEULSTORE=dir; pack tests clear it).
func newHarness(t *testing.T) *harness {
	data := t.TempDir()
	os.MkdirAll(filepath.Join(data, ".well-known", "geul"), 0o755)
	return &harness{t: t, data: data, conf: t.TempDir(), cwd: data, env: map[string]string{"GEULSTORE": "dir", "GEULPROXY": "off"}}
}
