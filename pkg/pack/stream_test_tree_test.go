//ff:func feature=pack type=writer control=sequence
//ff:what streamTestTree — every file under a set folder by path with its SHA-256
package pack

import (
	"crypto/sha256"
	"encoding/hex"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// streamTestTree walks out and maps each file's relative path to the hex
// SHA-256 of its bytes, so two sets compare byte for byte.
func streamTestTree(t *testing.T, out string) map[string]string {
	t.Helper()
	m := map[string]string{}
	err := filepath.WalkDir(out, func(p string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		b, err := os.ReadFile(p)
		sum := sha256.Sum256(b)
		rel, _ := filepath.Rel(out, p)
		m[rel] = hex.EncodeToString(sum[:])
		return err
	})
	if err != nil {
		t.Fatal(err)
	}
	return m
}
