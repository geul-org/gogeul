//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what TestStreamCloseMerge — a Close stopped in the middle of merging runs down (a run that cannot be read) keeps the merges it saved, and after the run is back and the stream resumed writes the files of an unstopped run
package pack

import (
	"maps"
	"os"
	"path/filepath"
	"testing"
)

func TestStreamCloseMerge(t *testing.T) {
	ref := streamCrashRef(t)
	for _, at := range []int{2, 6, 11} {
		out := filepath.Join(t.TempDir(), "set")
		s, _ := streamCrashStart(t, out, false)
		streamCrashRest(t, s)
		s.spill()
		run := filepath.Join(out, streamDir, s.st.Runs[at])
		good, _ := os.ReadFile(run)
		os.WriteFile(run, good[:hEntrySize+1], 0o644)
		if _, err := s.Close(); err == nil {
			t.Fatalf("run %d cut: want error", at)
		}
		b, _ := os.ReadFile(filepath.Join(out, streamDir, streamStateFile))
		if st := streamTestState(t, b); len(st.Close.Spent) != at/2*2 || st.Close.Phase != closePDone {
			t.Fatalf("run %d: saved %d spent runs, phase %d", at, len(st.Close.Spent), st.Close.Phase)
		}
		os.WriteFile(run, good, 0o644)
		if got := streamCrashResume(t, out); !maps.Equal(got, ref) {
			t.Fatalf("stop at run %d: files differ", at)
		}
	}
}
