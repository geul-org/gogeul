//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamOpenHPass — the pass reopens at its saved position (runs after the taken entries, head file cut to the saved heads, leaf and counts back); a missing run, a missing head file and a bad head count are errors
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamOpenHPass(t *testing.T) {
	s := streamTestNew(t, nil)
	s.st.Runs = streamTestRuns(t, s, []byte{1, 3}, []byte{2})
	s.hStart()
	heads := filepath.Join(s.out, streamDir, streamHeadsFile)
	os.WriteFile(heads, make([]byte, 3*hEntrySize), 0o644)
	c := &s.st.Close
	c.Taken, c.HLeaf, c.HKeys, c.HHeads = []int64{1, 0}, []streamKV{{Key: "k"}}, 5, 2
	p, err := s.openHPass()
	if err != nil {
		t.Fatal(err)
	}
	if fi, _ := os.Stat(heads); fi.Size() != 2*hEntrySize || p.rs[0].cur.key[0] != 3 || p.rs[0].taken != 1 || p.n != 5 || p.count != 2 || len(p.leaf) != 1 {
		t.Fatalf("size %d", fi.Size())
	}
	p.leaf[0].Key = "changed"
	if c.HLeaf[0].Key != "k" {
		t.Fatal("leaf shared with the state")
	}
	p.close()
	c.HHeads = -1
	if _, err := s.openHPass(); err == nil {
		t.Fatal("bad head count: want error")
	}
	os.Remove(heads)
	if _, err := s.openHPass(); err == nil {
		t.Fatal("missing head file: want error")
	}
	s.st.Runs = []string{"run-missing"}
	if _, err := s.openHPass(); err == nil {
		t.Fatal("missing run: want error")
	}
}
