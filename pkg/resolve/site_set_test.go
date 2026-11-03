//ff:func feature=resolve type=resolver control=sequence
//ff:what TestSiteSet — a set opens once, its absence is remembered, a broken one is not and heals when the bytes do
package resolve

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSiteSet(t *testing.T) {
	ctx := context.Background()
	root := packedRoot(t, formsTree(t), false, nil)
	s := localSite(root)
	a, err := s.set(ctx, "0x00000000")
	if err != nil || a == nil {
		t.Fatal(err)
	}
	if b, err := s.set(ctx, "0x00000000"); err != nil || b != a {
		t.Fatal("second open:", err)
	}
	if r, err := s.set(ctx, ""); err != nil || r != nil {
		t.Fatal(r, err)
	}
	if err := os.WriteFile(filepath.Join(root, ".well-known", "geul", "latest.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if r, err := s.set(ctx, ""); err != nil || r != nil {
		t.Fatal("absence is remembered:", r, err)
	}
	latest := filepath.Join(root, ".well-known", "geul", "0x00000002", "latest.txt")
	good, _ := os.ReadFile(latest)
	os.WriteFile(latest, []byte("x"), 0o644)
	if _, err := s.set(ctx, "0x00000002"); !errors.Is(err, pack.ErrCorrupt) {
		t.Fatal(err)
	}
	os.WriteFile(latest, good, 0o644)
	if r, err := s.set(ctx, "0x00000002"); err != nil || r == nil {
		t.Fatal(r, err)
	}
}
