//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestSiteFetch — a set's path, the loose file, nothing for an empty or absent place, and a broken latest.txt is an error and not ErrNotFound, file:// and https alike
package resolve

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSiteFetch(t *testing.T) {
	ctx := context.Background()
	tree := formsTree(t)
	addr := listPath(sampleKeyOf(7))
	empty := t.TempDir()
	os.MkdirAll(filepath.Join(empty, ".well-known", "geul"), 0o755)
	broken := packedRoot(t, tree, false, nil)
	os.WriteFile(filepath.Join(broken, ".well-known", "geul", "0x00000000", "latest.txt"), []byte("<html>not a set</html>"), 0o644)
	lost := packedRoot(t, tree, true, nil)
	os.WriteFile(filepath.Join(lost, ".well-known", "geul", "latest.txt"), []byte("3 97c284befd56c339bd9b2f40bb4990b3812375e89579130233afe05b6d0da9da\n"), 0o644)
	other := packedRoot(t, tree, false, nil)
	os.WriteFile(filepath.Join(other, ".well-known", "geul", "0x00000002", "latest.txt"), []byte("junk"), 0o644)
	sites := append(bothSites(t, empty), bothSites(t, filepath.Join(empty, "absent"))...)
	bad := append(bothSites(t, broken), bothSites(t, lost)...)
	good := append(bothSites(t, tree.Root), bothSites(t, other)...)
	for _, s := range sites {
		if _, err := s.Fetch(ctx, addr); !errors.Is(err, ErrNotFound) {
			t.Fatal(err)
		}
		if _, err := s.Fetch(ctx, issuersPath()); !errors.Is(err, ErrNotFound) {
			t.Fatal(err)
		}
	}
	for _, s := range bad {
		if _, err := s.Fetch(ctx, addr); !errors.Is(err, pack.ErrCorrupt) || errors.Is(err, ErrNotFound) {
			t.Fatal(err)
		}
	}
	for _, s := range good {
		if b, err := s.Fetch(ctx, addr); err != nil || len(b) == 0 {
			t.Fatal("another Issuer's broken set does not touch this one:", err)
		}
	}
	file := &site{dir: filepath.Join(broken, ".well-known", "geul", "0x00000000", "latest.txt")}
	if _, err := file.Fetch(ctx, issuersPath()); err == nil || errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
