//ff:func feature=resolve type=resolver control=iteration dimension=1
//ff:what TestResolverFetchOrder — GEULPROXY order over real places: a lacking place passes to the next, a broken one stops at `,` and passes at `|`, `off` stops
package resolve

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

func TestResolverFetchOrder(t *testing.T) {
	ctx := context.Background()
	tree := formsTree(t)
	good := "file://" + packedRoot(t, tree, false, &pack.Zstd{Level: 1})
	files := "file://" + tree.Root
	empty := "file://" + t.TempDir()
	root := packedRoot(t, tree, false, nil)
	if err := os.WriteFile(filepath.Join(root, ".well-known", "geul", "0x00000000", "latest.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	broken := "file://" + root
	k := Key{Issuer: []core.Word{0, 0}, Doc: []core.Word{0xffff, 0, 0, 7}}
	for proxy, source := range map[string]string{
		empty + "," + good:                good,
		good + "," + files:                good,
		files + "," + good:                files,
		empty + "," + empty + "," + files: files,
		broken + "|" + good:               good,
	} {
		c, err := ConfigFromEnv(func(k string) string { return map[string]string{"GEULPROXY": proxy}[k] })
		if err != nil {
			t.Fatal(err)
		}
		n, err := (&Resolver{Config: c}).Fetch(ctx, k)
		if err != nil || n.Source != source || n.Version != 2 {
			t.Fatalf("%s: %+v %v", proxy, n, err)
		}
	}
	c, _ := ConfigFromEnv(func(k string) string { return map[string]string{"GEULPROXY": broken + "," + good}[k] })
	if _, err := (&Resolver{Config: c}).Fetch(ctx, k); !errors.Is(err, pack.ErrCorrupt) {
		t.Fatal(err)
	}
	c, _ = ConfigFromEnv(func(k string) string { return map[string]string{"GEULPROXY": empty + ",off," + good}[k] })
	if _, err := (&Resolver{Config: c}).Fetch(ctx, k); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
}
