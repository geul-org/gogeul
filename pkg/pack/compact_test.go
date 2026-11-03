//ff:func feature=pack type=writer control=sequence
//ff:what TestCompact — after deltas, one generation with only new objects and single P and H layers serving the same paths; nothing to compact fails
package pack

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCompact(t *testing.T) {
	ctx := context.Background()
	root, out := built(t, small())
	h := "0x00000000/hashes/0x33BE25E73766F819226DD801DB778F61C503843D0E3A10CCB67258CD4108FE51.geul"
	os.Remove(filepath.Join(root, ".well-known", "geul", filepath.FromSlash(h)))
	if _, err := Build(ctx, root, out, small()); err != nil {
		t.Fatal(err)
	}
	before, _ := opened(t, out).Paths(ctx)
	if len(opened(t, out).m.hLayers) != 2 {
		t.Fatal("delta layer")
	}
	res, err := Compact(ctx, out, small())
	if err != nil || res.Generation != 3 {
		t.Fatal(res, err)
	}
	r := opened(t, out)
	after, _ := r.Paths(ctx)
	if !reflect.DeepEqual(before, after) || len(r.m.hLayers) != 1 || len(r.m.objects) != len(res.Objects) || r.Verify(ctx) != nil {
		t.Fatal(len(r.m.objects), len(res.Objects))
	}
	if _, err := Compact(ctx, t.TempDir(), small()); err == nil {
		t.Fatal("empty")
	}
	r.m.pLayers = nil
	os.Remove(filepath.Join(out, filepath.FromSlash(objectName(r.m.objects[len(r.m.objects)-1]))))
	if _, err := Compact(ctx, out, small()); err == nil {
		t.Fatal("paths")
	}
}
