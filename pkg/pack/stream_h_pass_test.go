//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamHPass — the runs merge into one layer of distinct keys from the saved position; a missing run, a cut run and a failing leaf write are errors
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamHPass(t *testing.T) {
	s := streamTestNew(t, nil)
	s.st.Runs = streamTestRuns(t, s, []byte{1, 3, 5, 7, 9}, []byte{2, 3, 6})
	s.hStart()
	ls, err := s.hPass()
	if err != nil || len(ls) != 1 || ls[0].count != 7 {
		t.Fatalf("%v %v", ls, err)
	}
	gone := streamTestNew(t, nil)
	gone.st.Runs = []string{"run-missing"}
	if _, err := gone.hPass(); err == nil {
		t.Fatal("missing run: want error")
	}
	cut := streamTestNew(t, nil)
	cut.st.Runs = streamTestRuns(t, cut, []byte{1, 2})
	cut.hStart()
	run := filepath.Join(cut.out, streamDir, cut.st.Runs[0])
	b, _ := os.ReadFile(run)
	os.WriteFile(run, b[:hEntrySize+1], 0o644)
	if _, err := cut.hPass(); err == nil {
		t.Fatal("cut run: want error")
	}
	bad := streamTestNew(t, failCodec{})
	bad.st.Runs = streamTestRuns(t, bad, []byte{1, 2, 3, 4, 5})
	bad.hStart()
	if _, err := bad.hPass(); err == nil {
		t.Fatal("failing leaf: want error")
	}
}
