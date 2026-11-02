//ff:func feature=pack type=resolver control=sequence
//ff:what TestFileStoreReadRange — exact ranges by pread; empty ranges need no file; missing is ErrNotFound; past the end is ErrCorrupt
package pack

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestFileStoreReadRange(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "f"), []byte("0123456789"), 0o644)
	s, ctx := FileStore{Root: dir}, context.Background()
	if b, err := s.ReadRange(ctx, "f", 7, 3); err != nil || string(b) != "789" {
		t.Fatal(err)
	}
	if b, err := s.ReadRange(ctx, "none", 0, 0); err != nil || len(b) != 0 {
		t.Fatal(err)
	}
	if _, err := s.ReadRange(ctx, "none", 0, 1); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := s.ReadRange(ctx, "f", 8, 3); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
