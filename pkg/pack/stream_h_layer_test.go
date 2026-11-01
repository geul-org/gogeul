//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamHLayer — runs merge into one layer counting distinct keys, a repeated key keeps its later record; no keys give no layer; a started pass goes on without spilling; a spill, a merge down, a pass start, a run, a merge or a leaf write that fails is an error
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamHLayer(t *testing.T) {
	s := streamTestNew(t, nil)
	var k1, k2 hEntry
	k1.key[0], k2.key[0] = 1, 2
	early, late := k1, k1
	early.loc.Off, late.loc.Off = 10, 20
	s.hbuf = []hEntry{late, k2}
	s.spill()
	s.hbuf = []hEntry{early}
	ls, err := s.hLayer()
	if err != nil || len(ls) != 1 || ls[0].count != 2 {
		t.Fatalf("%v %v", ls, err)
	}
	blk, _ := decodeBlock(mustUnzstd(t, ls[0].root))
	if blk.locs[0].Off != 20 {
		t.Fatalf("repeated key kept %d", blk.locs[0].Off)
	}
	if ls, err := streamTestNew(t, nil).hLayer(); err != nil || ls != nil {
		t.Fatal("no keys")
	}
	sp := streamTestNew(t, nil)
	sp.hbuf = []hEntry{k1}
	os.RemoveAll(filepath.Join(sp.out, streamDir))
	if _, err := sp.hLayer(); err == nil {
		t.Fatal("spill failure: want error")
	}
	gone := streamTestNew(t, nil)
	gone.st.Runs = []string{"run-missing"}
	if _, err := gone.hLayer(); err == nil {
		t.Fatal("missing run: want error")
	}
	cut := streamTestNew(t, nil)
	os.WriteFile(filepath.Join(cut.out, streamDir, "run-cut"), append(appendHEntry(nil, k1), 1), 0o644)
	cut.st.Runs = []string{"run-cut"}
	if _, err := cut.hLayer(); err == nil {
		t.Fatal("cut run: want error")
	}
	md := streamTestNew(t, nil)
	md.st.Runs, md.o.MergeRuns = []string{"run-a", "run-b", "run-c"}, 2
	if _, err := md.hLayer(); err == nil {
		t.Fatal("merge failure: want error")
	}
	hs := streamTestNew(t, nil)
	os.MkdirAll(filepath.Join(hs.out, streamDir, streamHeadsFile, "x"), 0o755)
	if _, err := hs.hLayer(); err == nil {
		t.Fatal("start failure: want error")
	}
	again := streamTestNew(t, nil)
	again.hbuf = []hEntry{k1}
	again.hStart()
	if ls, err := again.hLayer(); err != nil || ls != nil || len(again.st.Runs) != 0 {
		t.Fatal("pass already started: held keys are not spilled again")
	}
	bad := streamTestNew(t, failCodec{})
	for i := 0; i < 5; i++ {
		e := hEntry{}
		e.key[0] = byte(i)
		bad.hbuf = append(bad.hbuf, e)
	}
	if _, err := bad.hLayer(); err == nil {
		t.Fatal("leaf write failure: want error")
	}
}
