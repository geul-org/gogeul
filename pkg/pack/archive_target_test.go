//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestArchiveTarget — the holder; the Issuer's set; the top set; a new Issuer set; home; holder errors
package pack

import (
	"context"
	"testing"
)

func TestArchiveTarget(t *testing.T) {
	ctx := context.Background()
	a, wk := testArchive(t)
	i0 := a.sets[0]
	cases := []struct {
		name string
		want string
	}{
		{"issuers/0x00000001.geul", "0x00000000"},
		{"0x00000000/docs/0x0000000000000001/versions.geul", "0x00000000"},
		{"0x00000002/docs/0x0000000000000001/versions.geul", "0x00000002"},
		{"0x00000000/hashes/new.geul", "0x00000000"},
		{"other.geul", "0x00000000"},
	}
	for _, c := range cases {
		s, err := a.target(ctx, c.name)
		if err != nil || s.name != c.want {
			t.Fatal(c.name, s, err)
		}
	}
	top := a.set("")
	if s, _ := a.target(ctx, "0x00000003/docs/0x0000000000000001/versions.geul"); s != top {
		t.Fatal("top")
	}
	_ = wk
	i0.r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
	if _, err := a.target(ctx, "0x00000000/docs/0x0000000000000008/versions.geul"); err == nil {
		t.Fatal("broken")
	}
}
