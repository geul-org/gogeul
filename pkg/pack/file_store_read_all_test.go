//ff:func feature=pack type=resolver control=sequence
//ff:what TestFileStoreReadAll — a file's bytes; a missing file is ErrNotFound
package pack

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestFileStoreReadAll(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "latest.txt"), []byte("x"), 0o644)
	s := FileStore{Root: dir}
	if b, err := s.ReadAll(context.Background(), "latest.txt"); err != nil || string(b) != "x" {
		t.Fatal(err)
	}
	if _, err := s.ReadAll(context.Background(), "none"); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
