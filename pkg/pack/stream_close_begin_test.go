//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamCloseBegin — the first call spills the held keys and saves phase closeStarted with one block encoder; a later call saves nothing; a failing spill is an error
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamCloseBegin(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 3)
	if err := s.closeBegin(); err != nil || len(s.hbuf) != 0 || s.st.Close.Phase != closeStarted || s.o.Build.Codec.(*Zstd).Workers != 1 {
		t.Fatalf("%v", err)
	}
	b, _ := os.ReadFile(filepath.Join(s.out, streamDir, streamStateFile))
	if st := streamTestState(t, b); st.Close.Phase != closeStarted || st.Count != 6 || len(st.Data.Sum) == 0 {
		t.Fatalf("saved phase %d", st.Close.Phase)
	}
	os.RemoveAll(filepath.Join(s.out, streamDir))
	if err := s.closeBegin(); err != nil {
		t.Fatalf("again: %v", err)
	}
	sp := streamTestNew(t, nil)
	sp.hbuf = []hEntry{{}}
	os.RemoveAll(filepath.Join(sp.out, streamDir))
	if err := sp.closeBegin(); err == nil {
		t.Fatal("spill failure: want error")
	}
}
