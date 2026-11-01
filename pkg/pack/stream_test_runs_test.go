//ff:func feature=pack type=writer control=iteration dimension=1
//ff:what streamTestRuns — write sorted H runs into a stream's working folder: run i holds the given key bytes, each at offset key×10+i
package pack

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// streamTestRuns writes one run per keys list (sorted, hEntryLess) and
// returns the run names.
func streamTestRuns(t *testing.T, s *Stream, keys ...[]byte) []string {
	t.Helper()
	var names []string
	for i, ks := range keys {
		var es []hEntry
		for _, k := range ks {
			e := hEntry{loc: Loc{Off: uint64(k)*10 + uint64(i)}}
			e.key[0] = k
			es = append(es, e)
		}
		sort.Slice(es, func(a, b int) bool { return hEntryLess(es[a], es[b]) })
		var b []byte
		for _, e := range es {
			b = appendHEntry(b, e)
		}
		names = append(names, runName(100+i))
		if err := os.WriteFile(filepath.Join(s.out, streamDir, runName(100+i)), b, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return names
}
