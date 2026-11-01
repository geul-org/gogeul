//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamMergeDown — runs are merged MergeRuns at a time until at most MergeRuns remain, inputs kept as spent and the state saved; few runs are left alone; a failing merge or save and a Pace stop are errors
package pack

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamMergeDown(t *testing.T) {
	s := streamTestNew(t, nil)
	s.st.Runs = streamTestRuns(t, s, []byte{1}, []byte{2}, []byte{3}, []byte{4}, []byte{5})
	s.o.MergeRuns = 5
	if err := s.mergeDown(); err != nil || len(s.st.Runs) != 5 {
		t.Fatal("few runs")
	}
	s.o.MergeRuns = 2
	if err := s.mergeDown(); err != nil || len(s.st.Runs) != 2 || len(s.st.Close.Spent) != 6 || s.st.RunCount != 3 {
		t.Fatalf("runs %v spent %v %v", s.st.Runs, s.st.Close.Spent, err)
	}
	b, _ := os.ReadFile(filepath.Join(s.out, streamDir, streamStateFile))
	if st := streamTestState(t, b); len(st.Runs) != 2 || len(st.Close.Spent) != 6 {
		t.Fatal("not saved")
	}
	s.st.Runs = append([]string{"run-missing"}, s.st.Runs...)
	if err := s.mergeDown(); err == nil {
		t.Fatal("missing run: want error")
	}
	l := streamTestNew(t, nil)
	l.st.Runs = streamTestRuns(t, l, []byte{1}, []byte{2}, []byte{3})
	l.o.MergeRuns = 2
	os.MkdirAll(filepath.Join(l.out, streamDir, streamStateFile, "x"), 0o755)
	if err := l.mergeDown(); err == nil {
		t.Fatal("save failure: want error")
	}
	p := streamTestNew(t, nil)
	p.st.Runs = streamTestRuns(t, p, []byte{1}, []byte{2}, []byte{3})
	p.o.MergeRuns = 2
	p.o.Pace = func(string) error { return errors.New("stop") }
	if err := p.mergeDown(); !errors.Is(err, ErrStopped) || len(p.st.Runs) != 3 {
		t.Fatalf("Pace stop: %v", err)
	}
}
