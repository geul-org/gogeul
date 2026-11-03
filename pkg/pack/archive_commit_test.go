//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveCommit — staged writes become generations of the routed sets, a new Issuer set appears, a reopened archive reads the same; nothing staged is a no-op; failures keep the stage
package pack

import (
	"context"
	"errors"
	"io/fs"
	"testing"
)

func TestArchiveCommit(t *testing.T) {
	ctx := context.Background()
	a, wk := testArchive(t)
	if res, err := a.Commit(ctx); res != nil || err != nil {
		t.Fatal(res, err)
	}
	a.WriteFile("0x00000002/docs/0x0000000000000001/versions.geul", []byte("v"))
	a.Remove("issuers/0x00000001.geul")
	res, err := a.Commit(ctx)
	if err != nil || len(res) != 2 || a.Pending() {
		t.Fatal(res, err)
	}
	b, _ := OpenDir(ctx, wk, small())
	if got, err := b.ReadFile("0x00000002/docs/0x0000000000000001/versions.geul"); err != nil || string(got) != "v" || len(b.sets) != 2 {
		t.Fatal(err)
	}
	if _, err := b.ReadFile("issuers/0x00000001.geul"); !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	b.WriteFile("0x00000000/docs/0x0000000000000009/versions.geul", nil)
	b.find("0x00000000").r.m.named = append(b.find("0x00000000").r.m.named, namedEntry{"zzz", Loc{}})
	b.find("0x00000000").r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
	if _, err := b.Commit(ctx); err == nil || !b.Pending() {
		t.Fatal("routing", err)
	}
	c, _ := OpenDir(ctx, wk, small())
	c.WriteFile("0x00000003/docs/0x0000000000000001/versions.geul", []byte("v"))
	c.WriteFile("0x00000004/docs/0x0000000000000001/versions.geul", []byte("v"))
	c.o.Codec = fakeCodec{err: errors.New("x")}
	if _, err := c.Commit(ctx); err == nil || !c.Pending() {
		t.Fatal("write", err)
	}
}
