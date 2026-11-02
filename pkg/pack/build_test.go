//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestBuild — build, then local and HTTP Range reads equal the files with ≤ 2 Ranges per cold P read and ≤ 3 per H read (V6); a delta build hides changed and deleted paths (V4); a missing archive fails
package pack

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestBuild(t *testing.T) {
	ctx := context.Background()
	root, out := built(t, small())
	c := &counter{h: http.FileServer(http.Dir(out))}
	srv := httptest.NewServer(c)
	defer srv.Close()
	paths, _ := listTree(root)
	for _, p := range paths {
		r, err := Open(ctx, HTTPStore{Base: srv.URL}, Options{})
		if err != nil {
			t.Fatal(err)
		}
		c.ranges = 0
		b, err := r.Fetch(ctx, p)
		limit := map[int]int{spaceP: 2, spaceH: 3, spaceNamed: 1}[parsePath(p).space]
		if err != nil || !bytes.Equal(b, fixtureFile(t, p)) || c.ranges > limit {
			t.Fatal(p, err, c.ranges)
		}
	}
	if c.gets != 2*len(paths) {
		t.Fatal("GETs", c.gets)
	}
	base := filepath.Join(root, ".well-known", "geul")
	v1 := "0x0000/0x0000000000000845/0x00000001.geul"
	os.WriteFile(filepath.Join(base, "issuers.geul"), []byte("changed"), 0o644)
	os.Remove(filepath.Join(base, v1))
	res, err := Build(ctx, root, out, small())
	if err != nil || res.Generation != 2 || res.Records != 1 || res.Deleted != 1 {
		t.Fatal(res, err)
	}
	r := opened(t, out)
	if b, _ := r.Fetch(ctx, "issuers.geul"); string(b) != "changed" {
		t.Fatal(string(b))
	}
	if _, err := r.Fetch(ctx, v1); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if b, err := r.Fetch(ctx, "0x0000/0x0000000000000860/0x00000001.geul"); err != nil || !bytes.Equal(b, fixtureFile(t, "0x0000/0x0000000000000860/0x00000001.geul")) {
		t.Fatal(err)
	}
	if r.Verify(ctx) != nil || len(r.m.hLayers) != 1 {
		t.Fatal("verify")
	}
	if _, err := Build(ctx, t.TempDir(), out, small()); err == nil {
		t.Fatal("missing archive")
	}
}
