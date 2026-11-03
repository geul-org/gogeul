//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamCloseWidth — a merge width that changes between merges writes the same files; Close's peak heap follows the width, not the number of runs
package pack

import (
	"maps"
	"path/filepath"
	"testing"
)

func TestStreamCloseWidth(t *testing.T) {
	ref := streamCrashRef(t)
	out := filepath.Join(t.TempDir(), "set")
	s, _ := streamCrashStart(t, out, false)
	streamCrashRest(t, s)
	s.o.MergeRuns = 8
	asked := 0
	s.o.MergeWidth = func() int {
		asked++
		return 1 + asked%5
	}
	if _, err := s.Close(); err != nil {
		t.Fatal(err)
	}
	if !maps.Equal(streamTestTree(t, out), ref) || asked < 2 {
		t.Fatalf("files differ, width asked %d times", asked)
	}
	o := StreamOptions{Build: BuildOptions{Codec: Stored{}}, Issuers: []string{"0x00000001"}, RunEntries: 2, Created: 1_800_000_000, MergeWidth: func() int { return 4 }}
	n, err := NewStream(filepath.Join(t.TempDir(), "set"), o)
	if err != nil {
		t.Fatal(err)
	}
	streamTestPut(t, n, 1, 1000)
	var cerr error
	peak := peakHeap(func() { _, cerr = n.Close() })
	t.Logf("1000 runs, width 4: Close peak heap %.1f MiB", float64(peak)/(1<<20))
	if cerr != nil || peak > 4*MergeRunMemory+16<<20 {
		t.Fatalf("peak %d: %v", peak, cerr)
	}
}
