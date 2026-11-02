//ff:func feature=resolve type=helper control=iteration dimension=1
//ff:what formsOf — test fixture: the same tree as files, raw and zstd pack sets (per Issuer and at the top), each as a file:// source, a data folder and an https source
package resolve

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

// formsOf builds every storage form of tree.
func formsOf(t *testing.T, tree Dir) []form {
	t.Helper()
	files, _ := served(t, tree.Root)
	out := []form{
		{"files file://", Dir{Root: tree.Root, Store: localSite(tree.Root)}},
		{"files data folder", tree},
		{"files https", Dir{Store: files}},
	}
	for _, p := range []struct {
		name   string
		single bool
		codec  pack.Codec
	}{
		{"pack per Issuer", false, nil},
		{"pack top", true, nil},
		{"zstd per Issuer", false, &pack.Zstd{Level: 1}},
		{"zstd top", true, &pack.Zstd{Level: 1}},
	} {
		root := packedRoot(t, tree, p.single, p.codec)
		a, err := pack.OpenDir(context.Background(), filepath.Join(root, ".well-known", "geul"), pack.BuildOptions{})
		if err != nil {
			t.Fatal(err)
		}
		s, _ := served(t, root)
		out = append(out,
			form{p.name + " file://", Dir{Root: root, Store: localSite(root)}},
			form{p.name + " data folder", Dir{Root: root, Store: a}},
			form{p.name + " https", Dir{Store: s}})
	}
	return out
}
