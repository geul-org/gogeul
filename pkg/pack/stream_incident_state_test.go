//ff:func feature=pack type=writer control=sequence
//ff:what streamIncidentState — whether a stopped stream's folder is in the 2026-11-01 state: hash-less checkpoint with a data pack and an index object open, data.part gone, a sealed data pack and index object the checkpoint does not know, index.part present, a stray run
package pack

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// streamIncidentState checks out against the checkpoint bytes old.
func streamIncidentState(t *testing.T, out string, old []byte) bool {
	t.Helper()
	r := &Stream{out: out, st: streamTestState(t, old)}
	datas, _ := r.unknownObjects(kindData)
	indexes, _ := r.unknownObjects(kindIndex)
	_, dataErr := os.Stat(filepath.Join(out, streamDir, streamDataPart))
	_, indexErr := os.Stat(filepath.Join(out, streamDir, streamIndexPart))
	ents, _ := os.ReadDir(filepath.Join(out, streamDir))
	stray := slices.ContainsFunc(ents, func(e os.DirEntry) bool {
		return strings.HasPrefix(e.Name(), "run-") && !slices.Contains(r.st.Runs, e.Name())
	})
	return len(r.st.Data.Sum) == 0 && r.st.Data.Num >= 0 && r.st.Index.Num >= 0 && len(datas) > 0 && dataErr != nil && len(indexes) > 0 && indexErr == nil && stray
}
