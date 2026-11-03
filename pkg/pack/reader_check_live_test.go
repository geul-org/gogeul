//ff:func feature=pack type=resolver control=sequence
//ff:what TestReaderCheckLive — a live key stops the walk; a marker, a staged removal and a key hidden by a newer layer do not; a read error comes back
package pack

import (
	"context"
	"errors"
	"os"
	"testing"
)

func TestReaderCheckLive(t *testing.T) {
	ctx := context.Background()
	a, _ := testArchive(t)
	s := a.sets[0]
	kr, _ := s.r.dirRange("0x00000000/docs/0x0000000000000008")
	k, p := kr.lo, kr.path(kr.lo)
	if err := s.r.checkLive(ctx, kr, nil, k, Loc{}); !errors.Is(err, errLiveKey) {
		t.Fatal(err)
	}
	if err := s.r.checkLive(ctx, kr, nil, k, Loc{Enc: encGone}); err != nil {
		t.Fatal(err)
	}
	if err := s.r.checkLive(ctx, kr, map[string]bool{p: true}, k, Loc{}); err != nil {
		t.Fatal(err)
	}
	if err := s.r.checkLive(ctx, kr, nil, pKey(0, 0, 0), Loc{}); err != nil {
		t.Fatal("a key the index does not have:", err)
	}
	s.r.blocks = map[Loc]block{}
	os.Remove(objPath(t, s.r, s.dir, kindIndex))
	if err := s.r.checkLive(ctx, kr, nil, k, Loc{}); err == nil || errors.Is(err, errLiveKey) {
		t.Fatal("broken index:", err)
	}
}
