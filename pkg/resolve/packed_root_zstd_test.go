//ff:func feature=resolve type=helper control=sequence
//ff:what TestPackedRootZstd — the zstd form of the fixture really is compressed, so the form table reads compressed records
package resolve

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestPackedRootZstd(t *testing.T) {
	tree := formsTree(t)
	raw, zstd := treeSize(t, packedRoot(t, tree, false, nil)), treeSize(t, packedRoot(t, tree, false, &pack.Zstd{Level: 1}))
	if zstd >= raw {
		t.Fatalf("zstd %d bytes, raw %d", zstd, raw)
	}
}
