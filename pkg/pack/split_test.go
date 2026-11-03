//ff:func feature=pack type=writer control=sequence
//ff:what TestSplit — one set per Issuer that reads back every file; a missing archive, unreadable files and write errors fail
package pack

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestSplit(t *testing.T) {
	ctx := context.Background()
	root := fixture(t)
	wk := filepath.Join(t.TempDir(), "wk")
	rs, err := Split(ctx, root, wk, small())
	if err != nil || len(rs) != 1 || !rs[0].Changed {
		t.Fatal(rs, err)
	}
	a, _ := OpenDir(ctx, wk, small())
	paths, _ := a.Paths()
	if len(paths) != 18 || len(a.sets) != 1 || a.sets[0].name != "0x00000000" {
		t.Fatal(paths)
	}
	if _, err := Split(ctx, t.TempDir(), wk, small()); err == nil {
		t.Fatal("missing")
	}
	os.Chmod(filepath.Join(root, ".well-known", "geul", "0x00000000", "codecs", "0x000003c0", "0x00000001.geul"), 0)
	defer os.Chmod(filepath.Join(root, ".well-known", "geul", "0x00000000", "codecs", "0x000003c0", "0x00000001.geul"), 0o644)
	if _, err := Split(ctx, root, filepath.Join(t.TempDir(), "x"), small()); err == nil {
		t.Fatal("unreadable")
	}
	os.Chmod(filepath.Join(root, ".well-known", "geul", "0x00000000", "codecs", "0x000003c0", "0x00000001.geul"), 0o644)
	if _, err := Split(ctx, root, filepath.Join(t.TempDir(), "y"), BuildOptions{Codec: fakeCodec{err: os.ErrInvalid}}); err == nil {
		t.Fatal("write")
	}
}
