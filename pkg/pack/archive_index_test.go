//ff:func feature=pack type=resolver control=sequence
//ff:what TestArchiveIndex — sorted union of set paths with staged puts and removals; built once; set errors fail
package pack

import (
	"testing"
)

func TestArchiveIndex(t *testing.T) {
	a, _ := testArchive(t)
	a.sets = append(a.sets, &packSet{name: "0x00000009"})
	a.puts["issuers/0x00000002.geul"] = nil
	a.dels["issuers/0x00000001.geul"] = true
	idx, err := a.index()
	if err != nil || len(idx) != 18 || idx[len(idx)-1] != "issuers/0x00000002.geul" || idx[len(idx)-2] != "0x00000000/hashes/0xFD6C8B66528778724A2ACDE087CF312F9ABCFC402F8C9D2BB3A177E1CBD683B6.geul" || idx[0] != "0x00000000/codecs/0x00000420/0x00000001.geul" {
		t.Fatal(idx, err)
	}
	if again, _ := a.index(); &again[0] != &idx[0] {
		t.Fatal("rebuilt")
	}
	b, _ := testArchive(t)
	b.sets[0].r.m.pLayers = []layer{{blk: block{kind: blkPInner, keys: []string{"a"}, locs: []Loc{{Obj: 99}}}}}
	if _, err := b.index(); err == nil {
		t.Fatal("broken")
	}
}
