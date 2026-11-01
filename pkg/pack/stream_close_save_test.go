//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamCloseSave — the open objects are synced and the state saved with their hashes; a data or index object that cannot sync is an error
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamCloseSave(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	if err := s.closeSave(); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(filepath.Join(s.out, streamDir, streamStateFile))
	if st := streamTestState(t, b); len(st.Data.Sum) == 0 || len(st.Index.Sum) == 0 || st.Count != 12 {
		t.Fatal("hashes not saved")
	}
	s.index.Close()
	if err := s.closeSave(); err == nil {
		t.Fatal("index sync failure: want error")
	}
	s.data.Close()
	if err := s.closeSave(); err == nil {
		t.Fatal("data sync failure: want error")
	}
}
