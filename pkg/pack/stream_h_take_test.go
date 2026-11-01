//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamHTake — a repeated key keeps its last location, a full leaf is written before the next key, the position is saved every saveLeaves leaves; a failing leaf write or save is an error
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamHTake(t *testing.T) {
	s, p := streamTestHPass(t, nil)
	s.o.saveLeaves = 2
	entry := func(k byte, off uint64) hEntry { e := hEntry{loc: Loc{Off: off}}; e.key[0] = k; return e }
	for i, e := range []hEntry{entry(1, 10), entry(1, 20), entry(2, 0), entry(3, 0), entry(4, 0), entry(5, 0)} {
		if err := s.hTake(p, e); err != nil {
			t.Fatalf("entry %d: %v", i, err)
		}
	}
	if p.n != 5 || p.leaf[0].Key[0] != 5 {
		t.Fatalf("n %d leaf %v", p.n, p.leaf)
	}
	p.w.Flush()
	heads, _ := readHeads(p.heads, 0, 1)
	if p.count != 1 || heads[0].Key[0] != 1 {
		t.Fatalf("written %d", p.count)
	}
	for k := byte(6); k < 10; k++ {
		s.hTake(p, entry(k, 0))
	}
	if s.st.Close.HHeads != 2 || p.since != 0 {
		t.Fatalf("saved heads %d", s.st.Close.HHeads)
	}
	f, fp := streamTestHPass(t, failCodec{})
	for k := byte(1); k <= 4; k++ {
		f.hTake(fp, entry(k, 0))
	}
	if err := f.hTake(fp, entry(5, 0)); err == nil {
		t.Fatal("failing leaf write: want error")
	}
	l, lp := streamTestHPass(t, nil)
	l.o.saveLeaves = 1
	os.Remove(filepath.Join(l.out, streamDir, streamStateFile))
	os.MkdirAll(filepath.Join(l.out, streamDir, streamStateFile, "x"), 0o755)
	for k := byte(1); k <= 4; k++ {
		l.hTake(lp, entry(k, 0))
	}
	if err := l.hTake(lp, entry(5, 0)); err == nil {
		t.Fatal("failing save: want error")
	}
}
