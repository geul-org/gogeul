//ff:func feature=pack type=resolver control=sequence
//ff:what TestOpenDir — sets open at their latest generation; an empty folder is an empty archive; broken sets and folders fail
package pack

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenDir(t *testing.T) {
	ctx := context.Background()
	a, wk := testArchive(t)
	if len(a.sets) != 1 || a.sets[0].r.Generation() != 1 || a.dir != wk {
		t.Fatal(a.sets)
	}
	if a, err := OpenDir(ctx, t.TempDir(), BuildOptions{}); err != nil || len(a.sets) != 0 {
		t.Fatal(err)
	}
	os.WriteFile(filepath.Join(wk, "0x00000000", latestFile), []byte("x"), 0o644)
	if _, err := OpenDir(ctx, wk, BuildOptions{}); err == nil {
		t.Fatal("broken set")
	}
	f := filepath.Join(t.TempDir(), "f")
	os.WriteFile(f, nil, 0o644)
	if _, err := OpenDir(ctx, f, BuildOptions{}); err == nil {
		t.Fatal("not a directory")
	}
}
