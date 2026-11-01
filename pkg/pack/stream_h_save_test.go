//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamHSave — the pass position (taken per run, leaf, keys, heads) is saved with the head file flushed; a failing flush or head sync is an error
package pack

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamHSave(t *testing.T) {
	s, p := streamTestHPass(t, nil, []byte{1, 2}, []byte{3})
	nextMerged(p.rs)
	p.leaf, p.n, p.count, p.since = []streamKV{{Key: "k"}}, 4, 0, 3
	p.w.WriteString("x")
	if err := s.hSave(p); err != nil || p.since != 0 {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(filepath.Join(s.out, streamDir, streamStateFile))
	st := streamTestState(t, b)
	if st.Close.Taken[0] != 1 || st.Close.Taken[1] != 0 || st.Close.HKeys != 4 || len(st.Close.HLeaf) != 1 {
		t.Fatalf("%+v", st.Close)
	}
	if fi, _ := os.Stat(filepath.Join(s.out, streamDir, streamHeadsFile)); fi.Size() != 1 {
		t.Fatal("head file not flushed")
	}
	full, _ := os.OpenFile("/dev/full", os.O_WRONLY, 0)
	defer full.Close()
	w := p.w
	p.w = bufio.NewWriter(full)
	p.w.WriteString("x")
	if err := s.hSave(p); err == nil {
		t.Fatal("failing flush: want error")
	}
	p.w = w
	p.heads.Close()
	if err := s.hSave(p); err == nil {
		t.Fatal("failing head sync: want error")
	}
}
