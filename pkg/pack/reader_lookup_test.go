//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderLookup — slot 0 is versions.geul; unknown Issuers and documents are NotFound; a broken record location fails
package pack

import (
	"bytes"
	"context"
	"errors"
	"testing"
)

func TestReaderLookup(t *testing.T) {
	ctx := context.Background()
	_, out := built(t, small())
	r := opened(t, out)
	b, err := r.Lookup(ctx, "0x0000", 0x845, 0)
	if err != nil || !bytes.Equal(b, fixtureFile(t, "0x0000/0x0000000000000845/versions.geul")) {
		t.Fatal(err)
	}
	if _, err := r.Lookup(ctx, "0x0001", 0x845, 0); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	if _, err := r.Lookup(ctx, "0x0000", 0x846, 0); !errors.Is(err, ErrNotFound) {
		t.Fatal(err)
	}
	r.m.pLayers = []layer{{blk: block{kind: blkPLeaf, keys: []string{pKey(0, 1, 0)}, locs: []Loc{{Obj: 99}}}}}
	if _, err := r.Lookup(ctx, "0x0000", 1, 0); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
