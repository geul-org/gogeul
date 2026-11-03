//ff:func feature=resolve type=resolver control=iteration dimension=2
//ff:what TestSiteFetchMixed — one tree keeping an Issuer in a set beside another as files, or a set beside loose documents of its own Issuer, reads whole; the set wins a path both have
package resolve

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

func TestSiteFetchMixed(t *testing.T) {
	ctx := context.Background()
	tree := formsTree(t)
	seven := versionPath(sampleKeyOf(7), 1)
	other := "0x00000002/docs/0x0000000000000005"
	sum := func(is, p string) string {
		b, err := tree.Fetch(ctx, p)
		if err != nil {
			t.Fatal(p, err)
		}
		ws, _ := core.ReadWords(b)
		h, _ := core.ContentHash(ws)
		return hashPath(is, h)
	}
	paths := []string{issuersPath(), seven, listPath(sampleKeyOf(7)), other + "/versions.geul", other + "/0x00000001.geul", sum(geulIssuer, seven), sum("0x00000002", other+"/0x00000001.geul")}
	copyOut := func(root, p string) {
		b, err := tree.Fetch(ctx, p)
		if err != nil {
			t.Fatal(p, err)
		}
		if err := (Dir{Root: root}).write(p, b); err != nil {
			t.Fatal(err)
		}
	}
	// 0x0000 as a set, 0x0001 as files: the one set is rebuilt without 0x0001.
	mixed := packedRoot(t, tree, false, nil)
	wk := Dir{Root: mixed}.file("")
	if err := os.RemoveAll(filepath.Join(wk, "0x00000002")); err != nil {
		t.Fatal(err)
	}
	copyOut(mixed, other+"/versions.geul")
	copyOut(mixed, other+"/0x00000001.geul")
	copyOut(mixed, sum("0x00000002", other+"/0x00000001.geul"))
	remote, meter := served(t, mixed)
	a, err := pack.OpenDir(ctx, wk, pack.BuildOptions{})
	if err != nil {
		t.Fatal(err)
	}
	for name, s := range map[string]Source{"file://": localSite(mixed), "https": remote, "data folder": Dir{Root: mixed, Store: a}} {
		for _, p := range paths {
			want, _ := tree.Fetch(ctx, p)
			if b, err := s.Fetch(ctx, p); err != nil || !bytes.Equal(b, want) {
				t.Fatalf("%s: %s: %v", name, p, err)
			}
		}
	}
	if seen, _ := meter.take(); len(seen) == 0 {
		t.Fatal("no requests")
	}
	// A loose copy of a path the set has loses to the set, in both.
	if err := (Dir{Root: mixed}).write(seven, []byte("loose")); err != nil {
		t.Fatal(err)
	}
	want, _ := tree.Fetch(ctx, seven)
	for name, s := range map[string]Source{"file://": localSite(mixed), "https": remote} {
		if b, err := s.Fetch(ctx, seven); err != nil || !bytes.Equal(b, want) {
			t.Fatalf("%s: %q %v", name, b, err)
		}
	}
	// A top set beside an Issuer of files.
	top := packedRoot(t, tree, true, nil)
	copyOut(top, listPath(sampleKeyOf(7)))
	if err := (Dir{Root: top}).write("0x00000003/docs/0x0000000000000001/versions.geul", []byte("x")); err != nil {
		t.Fatal(err)
	}
	remote, _ = served(t, top)
	for name, s := range map[string]Source{"file://": localSite(top), "https": remote} {
		if b, err := s.Fetch(ctx, "0x00000003/docs/0x0000000000000001/versions.geul"); err != nil || string(b) != "x" {
			t.Fatalf("%s: %q %v", name, b, err)
		}
	}
}
