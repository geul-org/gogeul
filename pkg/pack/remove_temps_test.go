//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestRemoveTemps — temp files go and other files stay; an unreadable folder and an unremovable temp fail
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveTemps(t *testing.T) {
	dir := t.TempDir()
	for _, n := range []string{".tmp-1", ".tmp-2", "latest.txt"} {
		os.WriteFile(filepath.Join(dir, n), nil, 0o644)
	}
	if err := removeTemps(dir); err != nil {
		t.Fatal(err)
	}
	for n, want := range map[string]bool{".tmp-1": false, ".tmp-2": false, "latest.txt": true} {
		if _, err := os.Stat(filepath.Join(dir, n)); (err == nil) != want {
			t.Errorf("%s: kept %v", n, err == nil)
		}
	}
	if err := removeTemps(filepath.Join(dir, "none")); err == nil {
		t.Fatal("no folder: want error")
	}
	os.MkdirAll(filepath.Join(dir, ".tmp-3", "x"), 0o755)
	if err := removeTemps(dir); err == nil {
		t.Fatal("unremovable temp: want error")
	}
}
