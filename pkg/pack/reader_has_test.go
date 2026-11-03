//ff:func feature=pack type=resolver control=iteration dimension=1
//ff:what TestReaderHas — P, H and named paths without reading records; unknown Issuers and absent paths are false; index errors are returned
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderHas(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	for _, p := range []string{"0x00000000/docs/0x0000000000000845/versions.geul", "0x00000000/hashes/0x33BE25E73766F819226DD801DB778F61C503843D0E3A10CCB67258CD4108FE51.geul", "issuers/0x00000001.geul"} {
		if ok, err := r.Has(ctx, p); !ok || err != nil {
			t.Fatal(p, err)
		}
	}
	for _, p := range []string{"0x00000002/docs/0x0000000000000845/versions.geul", "0x00000000/docs/0x0000000000000846/versions.geul", "zz"} {
		if ok, err := r.Has(ctx, p); ok || err != nil {
			t.Fatal(p, err)
		}
	}
	r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
	if _, err := r.Has(ctx, "0x00000000/docs/0x0000000000000845/versions.geul"); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
