//ff:func feature=build type=builder control=iteration dimension=1
//ff:what TestTreeFrom — a file tree and its pack set (공표 주소 paths) bake to the same pages; a destination inside the source and an unreadable tree fail
package build

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestTreeFrom(t *testing.T) {
	root := t.TempDir()
	doc := core.PutWords(core.WrapDocument([]core.Word{core.TargetWord(core.KindMeta, core.MetaEnd)}))
	for _, p := range []string{"0x00000000/docs/0x0000000000000001/0x00000001.geul", "issuers/0x00000001.geul"} {
		f := filepath.Join(root, ".well-known", "geul", filepath.FromSlash(p))
		os.MkdirAll(filepath.Dir(f), 0o755)
		os.WriteFile(f, doc, 0o644)
	}
	wk := filepath.Join(t.TempDir(), ".well-known", "geul")
	if _, err := pack.Build(context.Background(), root, wk, pack.BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	a, _ := pack.OpenDir(context.Background(), wk, pack.BuildOptions{})
	dirOut, packOut := t.TempDir(), t.TempDir()
	c1, err1 := TreeFrom(resolve.Dir{Root: root}.Tree(), filepath.Join(root, ".well-known", "geul"), dirOut)
	c2, err2 := TreeFrom(a, wk, packOut)
	if err1 != nil || err2 != nil || c1 != c2 {
		t.Fatal(c1, c2, err1, err2)
	}
	page := filepath.Join(".well-known", "geul", "0x00000000", "docs", "0x0000000000000001", "0x00000001")
	want, _ := os.ReadFile(filepath.Join(dirOut, page))
	got, _ := os.ReadFile(filepath.Join(packOut, page))
	if string(got) != string(want) || len(got) == 0 {
		t.Fatal("pages differ")
	}
	if _, err := TreeFrom(a, wk, filepath.Dir(filepath.Dir(wk))); err == nil {
		t.Fatal("inside the source")
	}
	if _, err := TreeFrom(resolve.Dir{Root: t.TempDir()}.Tree(), "src", t.TempDir()); err == nil {
		t.Fatal("unreadable tree")
	}
}
