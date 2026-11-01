//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamCloseLayers — every step runs and saves once (a done step is skipped); a failing data pack, P layer, H layer, index object or save after any step is an error
package pack

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStreamCloseLayers(t *testing.T) {
	s := streamTestNew(t, nil)
	streamTestPut(t, s, 1, 6)
	s.closeBegin()
	if err := s.closeLayers(); err != nil || s.st.Close.Phase != closeHDone || s.st.Data.Num != -1 || s.st.Index.Num != -1 || len(s.st.Close.P) != 1 || len(s.st.Close.H) != 1 {
		t.Fatalf("%v", err)
	}
	if err := s.closeLayers(); err != nil {
		t.Fatalf("all done: %v", err)
	}
	closed := func() *os.File { f, _ := os.Open(os.DevNull); f.Close(); return f }
	lock := func(s *Stream) { p := filepath.Join(s.out, streamDir, streamStateFile); os.RemoveAll(p); os.MkdirAll(filepath.Join(p, "x"), 0o755) }
	cases := map[string]func(*Stream){
		"data pack":       func(s *Stream) { s.o.Build.Codec = failCodec{} },
		"save after data": func(s *Stream) { s.index = closed() },
		"P layer":         func(s *Stream) { s.finishData(); s.o.Build.Codec = failCodec{} },
		"save after P":    func(s *Stream) { s.finishData(); s.st.PHeads = nil; s.index = closed() },
		"H layer":         func(s *Stream) { s.finishData(); s.st.Close.Phase = closePDone; s.st.Runs = []string{"run-missing"} },
		"save after H":    func(s *Stream) { s.finishData(); s.spill(); s.hStart(); lock(s) },
		"index object":    func(s *Stream) { s.finishData(); s.st.Close.Phase = closeHDone; s.index.Close() },
		"save after index": func(s *Stream) {
			s.finishData()
			s.st.Close.Phase = closeHDone
			lock(s)
		},
	}
	for name, breakIt := range cases {
		b := streamTestNew(t, nil)
		streamTestPut(t, b, 1, 6)
		breakIt(b)
		if err := b.closeLayers(); err == nil {
			t.Errorf("%s: want error", name)
		}
	}
}
