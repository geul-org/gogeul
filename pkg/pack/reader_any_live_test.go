//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderAnyLive — true at a live key; false for an empty range and when every key is removed; a read error comes back
package pack

import (
	"context"
	"os"
	"testing"
)

func TestReaderAnyLive(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	s := a.sets[0]
	kr, _ := s.r.dirRange("0x0000/0x0000000000000008")
	if ok, err := s.r.anyLive(ctx, kr, nil); err != nil || !ok {
		t.Fatal(ok, err)
	}
	gone := map[string]bool{kr.path(pKey(0, 8, 0)): true, kr.path(pKey(0, 8, 1)): true}
	if ok, err := s.r.anyLive(ctx, kr, gone); err != nil || ok {
		t.Fatal(ok, err)
	}
	empty, _ := s.r.dirRange("0x0000/0x00000000000000ff")
	if ok, err := s.r.anyLive(ctx, empty, nil); err != nil || ok {
		t.Fatal(ok, err)
	}
	s.r.blocks = map[Loc]block{}
	os.Remove(objPath(t, s.r, s.dir, kindIndex))
	if _, err := s.r.anyLive(ctx, kr, nil); err == nil {
		t.Fatal("broken index: want error")
	}
}
