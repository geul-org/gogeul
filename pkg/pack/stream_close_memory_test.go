//ff:func feature=pack type=writer control=sequence
//ff:what TestStreamCloseMemory — Close's peak heap is bounded by MergeRuns read buffers, not by the number of runs: 100 and 1000 runs peak alike (gogeul 6c02e78 held 1 MiB per run), and fewer merge inputs peak lower
package pack

import (
	"path/filepath"
	"testing"
)

func TestStreamCloseMemory(t *testing.T) {
	peakOf := func(docs uint64, merge int) uint64 {
		o := StreamOptions{Build: BuildOptions{Codec: Stored{}}, Issuers: []string{"0x00000001"}, RunEntries: 2, Created: 1_800_000_000, MergeRuns: merge}
		s, err := NewStream(filepath.Join(t.TempDir(), "set"), o)
		if err != nil {
			t.Fatal(err)
		}
		streamTestPut(t, s, 1, docs)
		if len(s.st.Runs) < int(docs)-1 {
			t.Fatalf("%d runs for %d documents", len(s.st.Runs), docs)
		}
		var cerr error
		peak := peakHeap(func() { _, cerr = s.Close() })
		if cerr != nil {
			t.Fatal(cerr)
		}
		t.Logf("%d runs, MergeRuns %d: Close peak heap %.1f MiB", docs, s.o.MergeRuns, float64(peak)/(1<<20))
		return peak
	}
	const mib = 1 << 20
	few, many := peakOf(100, 0), peakOf(1000, 0)
	if many > DefaultMergeRuns*runBufSize+24*mib || many > few+16*mib {
		t.Fatalf("1000 runs peak %d, 100 runs %d", many, few)
	}
	if narrow := peakOf(300, 8); narrow > 8*runBufSize+16*mib {
		t.Fatalf("MergeRuns 8 peak %d", narrow)
	}
}
