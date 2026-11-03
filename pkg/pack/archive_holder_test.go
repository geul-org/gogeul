//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveHolder — the set that has the path; nil when none; index errors pass up
package pack

import (
	"context"
	"testing"
)

func TestArchiveHolder(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	if s, err := a.holder(ctx, "issuers/0x00000001.geul"); err != nil || s != a.sets[0] {
		t.Fatal(s, err)
	}
	if s, err := a.holder(ctx, "none"); err != nil || s != nil {
		t.Fatal(s, err)
	}
	a.sets[0].r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
	if _, err := a.holder(ctx, "0x00000000/docs/0x0000000000000008/versions.geul"); err == nil {
		t.Fatal("broken")
	}
}
