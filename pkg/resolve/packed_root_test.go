//ff:func feature=resolve type=helper control=selection
//ff:what packedRoot — test fixture: a file tree rewritten as pack sets (one per Issuer or one at the top) in a new data folder
package resolve

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

// packedRoot packs tree into a fresh root; c nil stores records raw.
func packedRoot(t *testing.T, tree Dir, single bool, c pack.Codec) string {
	t.Helper()
	root := t.TempDir()
	wk := filepath.Join(root, ".well-known", "geul")
	o := pack.BuildOptions{Codec: c, Verify: true}
	var err error
	switch {
	case single:
		_, err = pack.Build(context.Background(), tree.Root, wk, o)
	default:
		_, err = pack.Split(context.Background(), tree.Root, wk, o)
	}
	if err != nil {
		t.Fatal(err)
	}
	return root
}
