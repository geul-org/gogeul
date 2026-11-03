//ff:func feature=pack type=writer control=sequence
//ff:what TestArchiveBatches — puts grouped by target set (h/ with the commit's Issuer), removals to every set that may have them; routing errors fail
package pack

import (
	"context"
	"testing"
)

func TestArchiveBatches(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	a.WriteFile("0x00000002/docs/0x0000000000000001/versions.geul", []byte("v"))
	a.WriteFile("0x00000000/hashes/0xabababababababababababababababababababababababababababababababab.geul", []byte("h"))
	a.Remove("issuers/0x00000001.geul")
	bs, err := a.batches(ctx)
	if err != nil || len(bs) != 2 {
		t.Fatal(bs, err)
	}
	b1 := bs[a.find("0x00000002")]
	if len(b1.Put) != 1 || len(b1.Delete) != 0 {
		t.Fatal(b1)
	}
	b0 := bs[a.find("0x00000000")]
	if len(b0.Delete) != 1 || len(b0.Put) != 1 {
		t.Fatal(b0)
	}
	if got, err := b0.Get("0x00000000/hashes/0xabababababababababababababababababababababababababababababababab.geul"); err != nil || string(got) != "h" {
		t.Fatal(err)
	}
	a.find("0x00000000").r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{pKey(0, 0, 0)}, locs: []Loc{{Obj: 99}}}}}
	a.WriteFile("0x00000000/docs/0x0000000000000008/versions.geul", nil)
	if _, err := a.batches(ctx); err == nil {
		t.Fatal("only")
	}
	a.find("0x00000000").r.m.hLayers = []layer{{blk: block{kind: blkHInner, keys: []string{string(make([]byte, 32))}, locs: []Loc{{Obj: 99}}}}}
	a.puts = map[string][]byte{hPath("0x00000000", string(make([]byte, 32))): nil}
	if _, err := a.batches(ctx); err == nil {
		t.Fatal("target")
	}
}
