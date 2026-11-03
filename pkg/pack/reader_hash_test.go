//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderHash — an h/ document by hash; absent is NotFound; a record whose content hash differs is ErrCorrupt and ErrVerify; bad locations fail
package pack

import (
	"context"
	"errors"
	"testing"
)

func TestReaderHash(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	h, _ := contentHash(fixtureFile(t, "0x00000000/docs/0x0000000000000845/0x00000001.geul"))
	if _, err := r.Hash(ctx, "0x00000000", h); err != nil {
		t.Fatal(err)
	}
	if _, err := r.Hash(ctx, "0x00000000", [32]byte{1}); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := r.Hash(ctx, "0x00000002", h); !errors.Is(err, ErrNotFound) {
		t.Fatal("another Issuer's hashes/:", err)
	}
	other, _ := r.find(ctx, r.m.pLayers, pKey(0, 0x860, 1), "")
	r.m.hLayers = []layer{{blk: block{kind: blkHLeaf, keys: []string{string(h[:]), string([]byte{0xff, 31: 0})}, locs: []Loc{other, {Obj: 99}}}}}
	if _, err := r.Hash(ctx, "0x00000000", h); !errors.Is(err, ErrCorrupt) || !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
	if _, err := r.Hash(ctx, "0x00000000", [32]byte{0xff}); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	r.m.hLayers = []layer{{blk: block{kind: blkHInner, keys: []string{string(h[:])}, locs: []Loc{{Obj: 99}}}}}
	if _, err := r.Hash(ctx, "0x00000000", h); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
