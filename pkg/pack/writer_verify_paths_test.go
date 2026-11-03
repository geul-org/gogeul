//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestWriterVerifyPaths — all puts match and the Full set is complete; missing puts, wrong bytes, live deletions, index errors and a short set fail
package pack

import (
	"context"
	"crypto/sha256"
	"errors"
	"testing"
)

func TestWriterVerifyPaths(t *testing.T) {
	_, out := built(t, small())
	r := opened(t, out)
	w := &writer{ctx: context.Background(), sums: map[string][32]byte{}, b: Batch{Full: true}}
	paths, _ := r.Paths(w.ctx)
	for _, p := range paths {
		w.sums[p] = sha256.Sum256(fixtureFile(t, p))
	}
	if err := w.verifyPaths(r); err != nil {
		t.Fatal(err)
	}
	w.b.Full = false
	w.dels = []string{"gone.geul"}
	if err := w.verifyPaths(r); err != nil {
		t.Fatal(err)
	}
	w.b.Full = true
	w.sums["issuers/0x00000001.geul"] = [32]byte{}
	if err := w.verifyPaths(r); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
	w.sums = map[string][32]byte{"missing": {}}
	if err := w.verifyPaths(r); err == nil {
		t.Fatal("missing")
	}
	w.sums = map[string][32]byte{}
	if err := w.verifyPaths(r); !errors.Is(err, ErrCorrupt) {
		t.Fatal("short set", err)
	}
	w.dels = []string{"issuers/0x00000001.geul"}
	if err := w.verifyPaths(r); !errors.Is(err, ErrCorrupt) {
		t.Fatal("live deletion", err)
	}
	w.dels = nil
	r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	if err := w.verifyPaths(r); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
