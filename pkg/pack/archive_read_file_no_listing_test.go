//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestArchiveReadFileNoListing — absent paths and directory checks read a few blocks and never build the path list; a listing of the same set reads far more
package pack

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"strings"
	"testing"
)

func TestArchiveReadFileNoListing(t *testing.T) {
	const budget = 8
	ctx := context.Background()
	root := fixture(t)
	for i := range 200 {
		ballast(t, root, fmt.Sprintf("0x0000/0x00001000%08x/versions.geul", i))
	}
	out := t.TempDir()
	if _, err := Build(ctx, root, out, small()); err != nil {
		t.Fatal(err)
	}
	full := openMetered(t, out)
	if _, err := full.a.Paths(); err != nil || full.m.reads <= 10*budget {
		t.Fatalf("the meter does not see a listing: %d reads, %v", full.m.reads, err)
	}
	for _, name := range []string{"0x0000/0x0000ffff00000000/versions.geul", "0x0000/0x0000000000000008/0x00000063.geul", "0x0001/0x0000000000000001/versions.geul", "h/0x" + strings.Repeat("0", 64) + ".geul", "none"} {
		o := openMetered(t, out)
		if _, err := o.a.ReadFile(name); !errors.Is(err, fs.ErrNotExist) {
			t.Fatal(name, err)
		}
		if o.m.reads > budget || o.a.idx != nil {
			t.Fatalf("%s: %d reads, path list built %v", name, o.m.reads, o.a.idx != nil)
		}
	}
	o := openMetered(t, out)
	if !o.a.isDir("0x0000") || !o.a.isDir("h") || !o.a.isDir("0x0000/0x0000000000000008") || o.m.reads > budget || o.a.idx != nil {
		t.Fatalf("isDir: %d reads", o.m.reads)
	}
}
