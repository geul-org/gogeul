//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamCheckpoint — a closing stream saves only the meta; held keys spill, the state with meta is saved; a spill, a sync or a save that fails is an error
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamCheckpoint(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 2)
	if err := s.Checkpoint([]byte("m")); err != nil || len(s.hbuf) != 0 || len(s.st.Runs) != 1 {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(s.out, streamDir, streamStateFile)); err != nil {
		t.Fatal(err)
	}
	s.index, _ = os.Open(os.DevNull)
	s.index.Close()
	if err := s.Checkpoint(nil); err == nil {
		t.Fatal("index sync failure: want error")
	}
	s.index = nil
	s.data.Close()
	if err := s.Checkpoint(nil); err == nil {
		t.Fatal("data sync failure: want error")
	}
	s.data = nil
	streamTestPut(t, streamTestNew(t, nil), 1, 1)
	sp := streamTestNew(t, nil)
	sp.hbuf = []hEntry{{}}
	os.RemoveAll(filepath.Join(sp.out, streamDir))
	if err := sp.Checkpoint(nil); err == nil {
		t.Fatal("spill failure: want error")
	}
	if err := s.Checkpoint(nil); err != nil {
		t.Fatalf("all synced: %v", err)
	}
	cl := streamTestNew(t, nil)
	streamTestPut(t, cl, 1, 1)
	cl.st.Close.Phase = closeStarted
	cl.data.Close() // a closing checkpoint syncs nothing
	if err := cl.Checkpoint([]byte("later")); err != nil || len(cl.hbuf) == 0 {
		t.Fatalf("closing: %v", err)
	}
	b, _ := os.ReadFile(filepath.Join(cl.out, streamDir, streamStateFile))
	if st := streamTestState(t, b); string(st.Meta) != "later" || st.Close.Phase != closeStarted {
		t.Fatalf("closing meta %q", st.Meta)
	}
}
