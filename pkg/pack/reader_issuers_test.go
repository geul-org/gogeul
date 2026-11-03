//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderIssuers — Issuers with documents or named paths, in Dir order; Issuers with nothing left are skipped; walk errors fail
package pack

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestReaderIssuers(t *testing.T) {
	ctx := context.Background()
	m := &manifest{issuers: []string{"0x00000001", "0x00000000", "0x00000002", "0x00000003"},
		pLayers: []layer{{blk: block{kind: blkPLeaf, keys: []string{pKey(0, 1, 0), pKey(1, 1, 0)}, locs: []Loc{{Off: 64}, {Off: 64}}}}},
		named:   []namedEntry{{"0x00000002/codecs/0x00000420/0x00000001.geul", Loc{}}}}
	r := newReader(memStore{}, Options{}, m, "")
	got, err := r.Issuers(ctx)
	if err != nil || !reflect.DeepEqual(got, []string{"0x00000000", "0x00000001", "0x00000002"}) {
		t.Fatal(got, err)
	}
	r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 9}}}}}
	if _, err := r.Issuers(ctx); !errors.Is(err, ErrCorrupt) {
		t.Fatal(err)
	}
}
